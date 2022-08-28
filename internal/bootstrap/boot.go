package bootstrap

import (
	"fmt"
	"github-user-service/internal/adaptors/fetcher"
	"github-user-service/internal/domain/adaptors/logger"
	"github-user-service/internal/http"
	"github-user-service/internal/pkg/configs"
	"github-user-service/internal/pkg/logs"
	"github-user-service/internal/repository"
	services2 "github-user-service/internal/services"
	"github-user-service/internal/streaming"
	"os"
	"os/signal"

	inLog "log"
)

func Boot() {
	initConfigs()
	l := initLogger()

	fetcher := fetcher.NewFetcher()
	producer := streaming.NewSyncProducer()
	userService := services2.NewUserService(fetcher, producer)

	r := &http.Router{
		Config: &http.Config,
	}
	r.Init(l, userService)

	// exit signal channel
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	// exit all signal channel
	exitAll := make(chan bool, 1)

	// stop http router and metrics router with interrupt signal
	go func() {
		<-signals
		if err := r.Stop(); err != nil {
			inLog.Fatalln(fmt.Sprintf("failed to gracefully shutdown the server: %s", err))
		}
		exitAll <- true
	}()

	// start http router
	go func() {
		err := r.Start()
		if err != nil {
			inLog.Fatal(err)
		}
	}()
	<-exitAll
}

func initConfigs() {
	err := configs.Load(
		new(logs.LoggerConfig),
		new(http.RouterConf),
		new(streaming.KafkaConfig),
		new(repository.RepoConfig),
	)
	if err != nil {
		inLog.Fatal("error in loading configurations", err)
	}
}

func initLogger() logger.Logger {
	l, err := logs.NewLogger(logs.Config.Level)
	if err != nil {
		inLog.Fatalln("error loading new logger due to: ", err)
	}
	return l
}

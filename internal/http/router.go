package http

import (
	"context"
	"fmt"
	"github-user-service/internal/domain/adaptors/logger"
	"github-user-service/internal/domain/services"
	"github-user-service/internal/http/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

type Router struct {
	server *http.Server
	Config *RouterConf
	logger logger.Logger
}

func (r *Router) Init(l logger.Logger, userService services.UserService) {
	router := mux.NewRouter()
	r.logger = l

	r.server = &http.Server{
		Addr:         fmt.Sprintf(":%s", r.Config.Host),
		Handler:      router,
		ReadTimeout:  r.Config.Timeouts.Read,
		WriteTimeout: r.Config.Timeouts.Write,
	}

	router.Handle("/user/{user-name}", handlers.UserHandler{
		Log:         l,
		UserService: userService,
	})
}

func (r *Router) Start() error {
	r.logger.Info(fmt.Sprintf(`server starting on port :%s`, r.Config.Host))
	if err := r.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (r *Router) Stop() error {
	c, fn := context.WithTimeout(context.Background(), r.Config.Timeouts.ShoutDownWait)
	defer fn()
	r.logger.Info(fmt.Sprintf(`server shutting down on port :%s`, r.Config.Host))
	return r.server.Shutdown(c)
}

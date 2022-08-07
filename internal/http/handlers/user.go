package handlers

import (
	"context"
	"encoding/json"
	"github-user-service/internal/domain/adaptors/fetcher/user"
	"github-user-service/internal/domain/adaptors/logger"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	Log     logger.Logger
	Fetcher user.Fetcher
}

func (u UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars["user-name"]

	fetchedUser, err := u.Fetcher.Fetch(context.Background(), userName)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userByte, err := json.Marshal(&fetchedUser)
	if err != nil {
		return
	}

	_, err = w.Write(userByte)
	if err != nil {
		return
	}
}

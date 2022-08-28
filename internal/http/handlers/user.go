package handlers

import (
	"context"
	"encoding/json"
	"github-user-service/internal/domain/adaptors/logger"
	"github-user-service/internal/domain/services"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	Log         logger.Logger
	UserService services.UserService
}

func (u UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars["user-name"]

	fetchedUser, err := u.UserService.GetUser(context.Background(), userName)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	userByte, err := json.Marshal(&fetchedUser)
	if err != nil {
		return
	}

	_, err = w.Write(userByte)
	if err != nil {
		return
	}
}

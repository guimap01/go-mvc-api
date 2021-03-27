package controllers

import (
	"encoding/json"
	"go-mvc/src/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("UserId must be a number"))
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	userJson, _ := json.Marshal(user)

	resp.Write(userJson)

}

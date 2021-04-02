package controllers

import (
	"encoding/json"
	"go-mvc/src/services"
	"go-mvc/src/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "UserId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		resp.WriteHeader(apiErr.StatusCode)
		userErrJson, _ := json.Marshal(apiErr)
		resp.Write(userErrJson)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		userErrJson, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(userErrJson))
		return
	}

	userJson, _ := json.Marshal(user)

	resp.Write(userJson)

}

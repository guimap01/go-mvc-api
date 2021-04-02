package controllers

import (
	"go-mvc/src/services"
	"go-mvc/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "UserId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)

}

package controllers

import (
	"golang/micro/services"
	"golang/micro/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		// Handle the err and return to the client
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)
}
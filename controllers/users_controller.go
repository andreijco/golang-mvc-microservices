package controllers

import (
	"encoding/json"
	"golang/micro/services"
	"golang/micro/utils"
	"net/http"
	"strconv"
)


func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		// Handle the err and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(http.StatusNotFound)
		resp.Write(jsonValue)
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
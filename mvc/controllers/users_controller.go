package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adhikag24/golang-microservices/mvc/services"
	"github.com/adhikag24/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		// Just return the error the bad request to the client
		return
	}

	user, apiErr := services.GetUser(userId)

	if err != nil {
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(apiErr.Message))
		// Handle the err and return to the client
		return
	}

	// return user to client

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}

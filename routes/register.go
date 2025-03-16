package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/golang_sqlite_api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully registerd"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	var event models.Event
	event.ID = eventId
	event.UserID = userId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration for user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

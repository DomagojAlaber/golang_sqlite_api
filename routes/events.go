package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/golang_sqlite_api/models"
	"example.com/golang_sqlite_api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Print("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventByID(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event"})
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	var event *models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
	}

	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		fmt.Print("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		fmt.Print("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not update event"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		fmt.Print("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //checkIfEventExists

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		fmt.Print("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	err = event.Delete()

	if err != nil {
		fmt.Print("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}

func checkIfEventExists(eventId int64, context *gin.Context) (*models.Event, error) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return nil, err
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		fmt.Print("Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return nil, err
	}

	return event, nil
}

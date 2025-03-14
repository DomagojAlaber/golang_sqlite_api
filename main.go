package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/golang_sqlite_api/db"
	"example.com/golang_sqlite_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

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
	var event *models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("Error: ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
	}

	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

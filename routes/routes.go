package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	//events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	server.POST("/events", createEvent)

	server.PUT("/events/:id", updateEvent)

	server.DELETE("/events/:id", deleteEvent)

	//users
	server.POST("/signup", signupUser)
	server.POST("/login", loginUser)
}

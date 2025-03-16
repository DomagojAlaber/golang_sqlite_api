package routes

import (
	"example.com/golang_sqlite_api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register")
	authenticated.DELETE("/events/:id/register")

	//users
	server.POST("/signup", signupUser)
	server.POST("/login", loginUser)
}

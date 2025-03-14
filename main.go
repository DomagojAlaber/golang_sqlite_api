package main

import (
	"example.com/golang_sqlite_api/db"
	"example.com/golang_sqlite_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}

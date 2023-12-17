package main

import (
	"log"

	database "github.com/Chandan050/blogwebsite/database"
	"github.com/Chandan050/blogwebsite/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DbConnection()
	log.Printf("Successfully connected to database")

	app := gin.New()
	routes.Setup(app)
	app.Run(":8080")

}

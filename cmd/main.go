package main

import (
	"github.com/gin-gonic/gin"
	"vietime-backend/config"
	"vietime-backend/database/mongodb"
	"vietime-backend/internal/delivery/http/route"
)

func main() {
	// Get env config
	bootstrap.NewEnv()

	// Connect to database
	db := mongodb.NewDBConnection(bootstrap.E.MongoDBURI)

	router := gin.Default()
	route.Setup(db, router)
	router.Run(bootstrap.E.ServerAddress)
	println(db.Name())
}

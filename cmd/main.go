package main

import (
	"vietime-backend/config"
	"vietime-backend/database/mongodb"
)

func main() {
	// Get env config
	bootstrap.NewEnv()

	// Connect to database
	db := mongodb.NewDBConnection(bootstrap.E.MongoDBURI)

	println(db.Name())
}

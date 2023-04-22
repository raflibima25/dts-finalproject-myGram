package main

import (
	"final-project-mygram/database"
	"final-project-mygram/router"
)

func main() {

	database.StartDB()

	router.StartApp().Run(":8080")

}

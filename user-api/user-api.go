package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"database"
)

func main() {
	defer database.DB.Close()

	r := gin.Default()
	r.GET("/ping", getPing)
	r.Run() // listen and serve on 0.0.0.0:8080

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}
	log.Println("connected to DB")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}
}
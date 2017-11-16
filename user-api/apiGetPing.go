package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	// "database"
)

func getPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
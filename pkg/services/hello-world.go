package services

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func HelloWorld(c *gin.Context) {

	count := 0
	
	
	for i := 0; i < 1000; i++ {
		count++
	}

	c.JSON(http.StatusOK, gin.H{
		"message": count,
	  })
}

package services

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"html/template"
	"net/http"
	"time"
)

func HelloWorld(c *gin.Context) {

	count := 0

	for i := 0; i < 1000; i++ {
		count++
	}

	tmpl := template.Must(template.New("hello").Parse("<h1>Hello, World!</h1><br>"))
	tmpl.Execute(c.Writer, nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "message",
	})

	go expensiveOperation()
}

func expensiveOperation() {
	time.Sleep(10 * time.Second)
	fmt.Println("Expensive operation completed")
}

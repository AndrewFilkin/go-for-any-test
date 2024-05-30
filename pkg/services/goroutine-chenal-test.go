package services

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"html/template"
	"time"
)

func ChenalTest(c *gin.Context) {

	message := "Test Message"
	tmpl := template.Must(template.New("goroutine").Parse("<h1>" + message + "</h1><br>"))
	tmpl.Execute(c.Writer, nil)

	ch := make(chan string)
	ch1 := make(chan string)

	go func() {
		ch <- "Chenal 1"
		time.Sleep(2 * time.Second)
	}()

	go func() {
		ch1 <- "Chenal 1"
		time.Sleep(1 * time.Second)
	}()

	select {
	case rec1 := <-ch:
		fmt.Println(rec1)
	case rec2 := <-ch1:
		fmt.Println(rec2)
	}

	// go sendData(ch, "hello")
	// go receiveData(ch)

}

func sendData(ch chan string, message string) {
	time.Sleep(5 * time.Second)
	ch <- message
}

func receiveData(ch chan string) {
	fmt.Println(<-ch)
}

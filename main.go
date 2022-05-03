package main

import (
	"log"
	"net/http"

	// "github.com/go-fsnotify/fsnotify"

	"github.com/gin-gonic/gin"
)

// Book ...
type Book struct {
    Title  string
    Author string
}

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "books": books,
        })
    })

    log.Fatal(r.Run())
}
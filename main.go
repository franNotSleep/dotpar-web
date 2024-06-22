package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.LoadHTMLGlob("view/html/*")
	r.Static("/css", "./view/css")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "main site",
		})
	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func compile(content string) string {
	cmd := exec.Command("./dotpar")

	stdin, err := cmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, content)
	}()

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

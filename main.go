package main

import (
	"io"
	"log"
	"net/http"

	"os/exec"

	"github.com/gin-gonic/gin"
)

type CompileBody struct {
	Content string `json:"content"`
}

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.LoadHTMLGlob("view/html/*")
	r.Static("/css", "./view/css")
	r.Static("/js", "./view/js")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/compile", func(c *gin.Context) {
		body := CompileBody{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		compiledContent, err := compile(body.Content)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			log.Println(err)
		}

		response := map[string]interface{}{
			"compiled_content": compiledContent,
		}

		c.JSON(http.StatusOK, response)
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

func compile(content string) (string, error) {
	cmd := exec.Command("./dotpar")

	stdin, err := cmd.StdinPipe()

	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, content)
	}()

	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", nil
	}

	return string(out), nil
}

package main

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStaticFiles(t *testing.T) {
	router := gin.Default()
	router.Static("/img", "./static/img")
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("/newcake.png", "./static/img/cake.png")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

func TestReader(t *testing.T) {
	router := gin.Default()
	router.GET("/reader", func(c *gin.Context) {
		response, err := http.Get("http://localhost:8080/newcake.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8081")
}

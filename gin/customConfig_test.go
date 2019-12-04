package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestCustomHTTPConfig(t *testing.T) {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

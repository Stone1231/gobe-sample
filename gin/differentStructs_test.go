package main

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		//   ...
	}
}

func TestDifferentStructs(t *testing.T) {
	r := gin.Default()
	r.POST("/diff", SomeHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}

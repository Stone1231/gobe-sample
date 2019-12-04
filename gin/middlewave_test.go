package main

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"
)

func mw(c *gin.Context) {
	log.Printf("IP: %s \n", c.ClientIP())
}

func TestMiddleware(t *testing.T) {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", mw, endpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(mw)
	{
		authorized.GET("/login", endpoint)
		authorized.POST("/submit", endpoint)
		authorized.POST("/read", endpoint)

		// nested group
		testing := authorized.Group("group")
		testing.GET("/sub", endpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

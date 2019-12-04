package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func TestBindURI(t *testing.T) {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": user.Name, "uuid": user.ID})
	})
	route.Run(":8088")
}

package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func TestBindHeader(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		h := testHeader{}
		rate, err := strconv.Atoi(c.GetHeader("Rate"))
		if err != nil {
			c.JSON(200, err)
		}
		h.Rate = rate
		h.Domain = c.GetHeader("Domain")

		// if err := c.ShouldBindHeader(&h); err != nil {
		// 	c.JSON(200, err)
		// }

		fmt.Printf("%#v\n", h)
		c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})

	r.Run() //default 8080

	// client
	// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/
	// output
	// {"Domain":"music","Rate":300}
}

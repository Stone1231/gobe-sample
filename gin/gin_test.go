package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestQuickStart(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

// func aPIExamples(){
// 	// Creates a gin router with default middleware:
// 	// logger and recovery (crash-free) middleware
// 	router := gin.Default()

// 	router.GET("/someGet", getting)
// 	router.POST("/somePost", posting)
// 	router.PUT("/somePut", putting)
// 	router.DELETE("/someDelete", deleting)
// 	router.PATCH("/somePatch", patching)
// 	router.HEAD("/someHead", head)
// 	router.OPTIONS("/someOptions", options)

// 	// By default it serves on :8080 unless a
// 	// PORT environment variable was defined.
// 	router.Run()
// 	// router.Run(":3000") for a hard coded port
// }

func TestParametersInPath(t *testing.T) {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	c.FullPath() == "/user/:name/*action" // true
	// })

	router.Run(":8080")
}

func TestQuerystringParameters(t *testing.T) {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}

// Multipart/Urlencoded Form
func TestMultipartUrlencodedForm(t *testing.T) {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}

func TestQueryPostForm(t *testing.T) {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}

//Map as querystring or postform parameters
func TestMapParameters(t *testing.T) {

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}

//Upload files
func TestUploadFiles(t *testing.T) {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		err = c.SaveUploadedFile(file, dst+"/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}

//Multiple files
func TestUploadMultipleFiles(t *testing.T) {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	dst, err := os.Getwd()
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
			err = c.SaveUploadedFile(file, dst+"/"+file.Filename)
			if err != nil {
				log.Fatal(err)
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}

func endpoint(c *gin.Context) {
	get := c.Query("get")
	post := c.PostForm("post")
	log.Printf("get: %s ", get)
	log.Printf("post: %s", post)
	c.String(http.StatusOK, "ok!")
}

func TestGroupingRoutes(t *testing.T) {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", endpoint)
		v1.POST("/submit", endpoint)
		v1.POST("/read", endpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", endpoint)
		v2.POST("/submit", endpoint)
		v2.POST("/read", endpoint)
	}

	router.Run(":8080")
}

func TestSecureJSON(t *testing.T) {
	r := gin.Default()

	// You can also use your own secure json prefix
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/json", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// curl http://127.0.0.1:8080/JSONP?callback=x
func TestJsonp(t *testing.T) {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

	// client
	// curl http://127.0.0.1:8080/JSONP?callback=x
}

// curl http://127.0.0.1:8080/json
func TestAsciiJSON(t *testing.T) {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// curl http://127.0.0.1:8080/json
// curl http://127.0.0.1:8080/purejson
// func pureJSON() {
func TestPureJSON(t *testing.T) {
	r := gin.Default()

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

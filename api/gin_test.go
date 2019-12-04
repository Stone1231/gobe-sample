package main

import (
	"log"
	"net/http"
	"testing"
	"time"
	. "be-ex/models"
	. "github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
)

func postUser(c *gin.Context) {
	var model User

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.Birthday, _ = time.Parse("2006-01-02", model.BirthdayStr)

	db.First(&(model.Dept), model.DeptID)

	db.Where("ID in (?)", model.ProjIDs).Find(&model.Projs)

	if err = db.Create(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(model.Name)
	log.Println(model.ID)
	log.Println(model.Birthday)
	log.Println(model.Dept)
	log.Println(model.Projs)
	log.Println(model)

	c.JSON(http.StatusOK, model)
}

func putUser(c *gin.Context) {
	var model User

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// uri/:id
	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	model.Birthday, _ = time.Parse("2006-01-02", model.BirthdayStr)

	db.First(&(model.Dept), model.DeptID)

	db.First(&User{}, model.ID).Association("Projs").Clear() //ori clear
	db.Where("ID in (?)", model.ProjIDs).Find(&model.Projs)
	// db.Where("ID in (?)", model.ProjIDs).Find(&(model.Projs))

	// if err = db.Update(&model).Error; err != nil {  !!!ERROR!!!
	if err = db.Save(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// if err = db.Model(&model).
	// 	Association("Projs").
	// 	Replace(&(model.Projs)).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	log.Println(model.Name)
	log.Println(model.ID)
	log.Println(model.Birthday)
	log.Println(model.Dept)
	log.Println(model.Projs)
	log.Println(model)

	c.JSON(http.StatusOK, model)
}

func getUser(c *gin.Context) {
	var model User

	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	db.First(&model, model.ID).Related(&model.Dept, "Dept")
	// If the field name is same as the variable's type name, like above example, it could be omitted, like:
	// db.Model(&model).Related(&model.Dept)

	db.Model(&model).Related(&model.Projs, "Projs")
	// db.Preload("Projs").First(&model)
	model.BirthdayStr = model.Birthday.Format("2006-01-02")

	From(model.Projs).Select(func(c interface{}) interface{} {
		return c.(Proj).ID
	}).ToSlice(&(model.ProjIDs))

	c.JSON(http.StatusOK, model)
}

func getUserAll(c *gin.Context) {
	var list []User

	db.Find(&list)

	c.JSON(http.StatusOK, list)
}

func TestApiDept(t *testing.T) {
	router := gin.Default()
	router.POST("/dept", postDept)
	router.Run(":8080")
}

func postDept(c *gin.Context) {
	var model Dept
	if c.ShouldBindJSON(&model) != nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(model.Name)
		log.Println(model.ID)
	}
	c.String(200, "Success")
}

// {"id":0,"name":"test","hight":"11","dept":"1","projs":[2,3],"photo":"","birthday":"2019-11-01"}

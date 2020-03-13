package main

import (
	"log"
	"os"
	"testing"
	"time"

	. "be-ex/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB
var err error

func TestMain(m *testing.M) {
	log.Println("test start")
	dbName := "test.db"
	db, err = gorm.Open("sqlite3", dbName)
	db.AutoMigrate(&Dept{}, &Proj{}, &User{})
	if err != nil {
		panic("failed to connect database")
	}
	retCode := m.Run()
	db.Close()
	os.Remove(dbName)
	os.Exit(retCode)
}

func TestDeptInit(t *testing.T) {
	tx := db.Begin()

	// if err := tx.Delete(&Dept{}).Error; err != nil {
	// 	tx.Rollback()
	// }

	names := []string{"dept1", "dept2", "dept3"}
	count := len(names)

	for i := 0; i < count; i++ {
		model := &Dept{ID: uint(i + 1), Name: names[i]}
		if err = tx.Create(model).Error; err != nil {
			tx.Rollback()
		}
	}

	tx.Commit()
}

func TestProjInit(t *testing.T) {
	tx := db.Begin()

	// if err := tx.Delete(&Proj{}).Error; err != nil {
	// 	tx.Rollback()
	// }

	names := []string{"proj1", "proj2", "proj3"}
	count := len(names)

	for i := 0; i < count; i++ {
		model := &Proj{ID: uint(i + 1), Name: names[i]}
		if err = tx.Create(model).Error; err != nil {
			tx.Rollback()
		}
	}

	tx.Commit()
}

func TestUserInit(t *testing.T) {
	tx := db.Begin()

	// if err := tx.Delete(&User{}).Error; err != nil {
	// 	tx.Rollback()
	// }

	model := &User{}
	model.Name = "user1"
	model.Hight = 170
	model.Photo = ""
	model.Birthday = time.Date(1977, time.December, 31, 0, 0, 0, 0, time.UTC)

	var dept Dept
	tx.First(&dept, 1)
	model.Dept = dept

	var projs []Proj
	tx.Find(&projs)

	model.Projs = projs

	if err = tx.Create(model).Error; err != nil {
		tx.Rollback()
	}

	model2 := model
	var dept2 Dept

	tx.First(&dept2, 2)
	model2.ID = 0
	model2.Dept = dept2
	if err = tx.Create(model2).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()
}

func TestInit(t *testing.T) {
	TestDeptInit(t)
	TestProjInit(t)
	TestUserInit(t)
}

func TestUserDelete(t *testing.T) {
	TestInit(t)

	tx := db.Begin()
	var model User
	tx.First(&model, 1)
	// db.Delete(&model)

	// db.Delete(&User{}, "1")
	tx.Model(&model).Association("projs").Clear()
	tx.Delete(&model)

	tx.Commit()
}

func TestDeptDelete(t *testing.T) {
	TestInit(t)

	tx := db.Begin()

	var model Dept

	// maybe record not found will delete all dept
	if err := tx.First(&model, 1).Error; err != nil {
		return
	}

	tx.Model(&model).Related(&model.Users, "Users")
	// tx.First(&model, 1).Related(&model.Users, "Users")

	// delete user_proj
	for i := 0; i < len(model.Users); i++ {
		tx.Model(&(model.Users[i])).Association("projs").Clear()
	}

	tx.Delete(User{}, "dept_id = ?", &model.ID)

	// //delete all users!! not dept's users!
	// tx.Delete(&model.Users)

	// //just update dept's users.DeptId -> null
	// tx.Model(&model).Association("users").Clear()

	tx.Delete(&model)

	tx.Commit()
}

func TestProjDelete(t *testing.T) {
	TestInit(t)

	tx := db.Begin()
	var model Proj
	tx.First(&model, 1)

	tx.Model(&model).Association("users").Clear()
	tx.Delete(&model)

	tx.Commit()
}

func TestQuickStart(t *testing.T) {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}

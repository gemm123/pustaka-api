package main

import (
	"gemm123/pustaka-api/handler"
	"gemm123/pustaka-api/models"
	"gemm123/pustaka-api/repository"
	"gemm123/pustaka-api/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=gemmq123456 dbname=pustaka_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to database!")
	}

	db.AutoMigrate(&models.Book{})

	reposittory := repository.NewRepository(db)
	service := service.NewService(reposittory)

	book := models.Book{
		Title: "why we sleep",
		Price: 75000,
	}

	service.Create(book)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/books/:id", handler.GetBooksByID)
	router.GET("/books", handler.GetBooksQuery)
	router.POST("/book", handler.PostBook)

	router.Run()
}

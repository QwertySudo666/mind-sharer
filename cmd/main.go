package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mind-sharer/infrastructure/datastorage"
	"mind-sharer/infrastructure/web"
	"mind-sharer/internal/domain/models"
	"mind-sharer/usecases"
)

func main() {
	router := gin.Default()

	// Initialize the database connection
	dsn := "root:localhost@tcp(127.0.0.1:3306)/mind_sharer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the database
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	// Initialize the repository
	usersRepository := datastorage.NewUsersRepository(db)

	// Initialize the use case
	usersUseCase := usecases.NewUsersUseCase(usersRepository)

	usersHandler := web.NewUsersHandler(usersUseCase)

	usersHandler.RegisterUsersRoutes(router)

	router.Run(":8080")
}

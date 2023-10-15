package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mind-sharer/domain/models"
	"mind-sharer/persistence"
	"mind-sharer/usecases"
)

func Main() {
	router := gin.Default()

	// Initialize the database connection
	dsn := "root:localhost@tcp(127.0.0.1:3306)/mind_sharer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the database
	db.AutoMigrate(&models.User{})

	// Initialize the repository
	userRepo := persistence.NewUserRepository(db)

	// Initialize the use case
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Define API routes
	router.GET("api/users", func(context *gin.Context) {
		users, err := userUseCase.SearchUsers()
		if err != nil {
			context.JSON(404, err)
			return
		}
		context.IndentedJSON(200, users)
	})

	router.GET("api/users/:id", func(c *gin.Context) {
		userID := uuid.MustParse(c.Param("id"))
		user, err := userUseCase.GetUserByID(userID)
		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	})

	router.POST("api/users", func(c *gin.Context) {
		var user models.User
		if err := c.Bind(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		userID, err := userUseCase.RegisterUser(user)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(201, gin.H{"user_id": userID})
	})

	router.PUT("api/users", func(c *gin.Context) {
		var user models.User
		if err := c.Bind(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		userID, err := userUseCase.UpdateUser(user)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(201, gin.H{"user_id": userID})
	})

	router.DELETE("api/users/:id", func(context *gin.Context) {
		userID := uuid.MustParse(context.Param("id"))
		user, err := userUseCase.DeleteUser(userID)
		if err != nil {
			fmt.Println(err)
		}
		context.IndentedJSON(200, user)
	})

	// Start the server
	router.Run(":8080")
}

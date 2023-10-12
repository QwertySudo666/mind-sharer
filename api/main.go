package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mind-sharer/domain"
	"mind-sharer/persistence"
	"mind-sharer/usecases"
)

func Main() {
	r := gin.Default()

	// Initialize the database connection
	dsn := "root:localhost@tcp(127.0.0.1:3306)/mind_sharer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the database
	db.AutoMigrate(&domain.User{})

	// Initialize the repository
	userRepo := persistence.NewUserRepository(db)

	// Initialize the use case
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Define API routes
	r.POST("api/users", func(c *gin.Context) {
		var user domain.User
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

	r.GET("api/users/:id", func(c *gin.Context) {
		userID := uuid.MustParse(c.Param("id"))
		user, err := userUseCase.GetUserByID(userID)
		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	})

	// Start the server
	r.Run(":8080")
}

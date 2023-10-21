package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mind-sharer/internal/domain/models"
	"mind-sharer/usecases"
)

type UsersHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewUsersHandler(userUseCase *usecases.UserUseCase) *UsersHandler {
	return &UsersHandler{userUseCase}
}

func (u *UsersHandler) FetchAllUsers(ctx *gin.Context) {
	users, err := u.userUseCase.SearchUsers()
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	ctx.IndentedJSON(200, users)
}

func (u *UsersHandler) FetchUserById(ctx *gin.Context) {
	userID := uuid.MustParse(ctx.Param("id"))
	user, err := u.userUseCase.GetUserByID(userID)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, user)
}

func (u *UsersHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	userID, err := u.userUseCase.RegisterUser(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(201, gin.H{"user_id": userID})
}

func (u *UsersHandler) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	userID, err := u.userUseCase.UpdateUser(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(201, gin.H{"user_id": userID})
}

func (u *UsersHandler) DeleteUserById(ctx *gin.Context) {
	userID := uuid.MustParse(ctx.Param("id"))
	user, err := u.userUseCase.DeleteUser(userID)
	if err != nil {
		fmt.Println(err)
	}
	ctx.IndentedJSON(200, user)
}

func (u *UsersHandler) RegisterUsersRoutes(router *gin.Engine) {
	router.GET("api/users", u.FetchAllUsers)
	router.GET("api/users/:id", u.FetchUserById)
	router.POST("api/users", u.CreateUser)
	router.PUT("api/users", u.UpdateUser)
	router.DELETE("api/users/:id", u.DeleteUserById)
}

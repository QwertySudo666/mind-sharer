package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func configureUsersRouter() {

	router := gin.Default()
	router.GET("api/users", getUsers)
	router.GET("api/users/:id", fetchUser)
	router.POST("api/users", saveUser)
	//router.PUT("api/users/:id", updateUser)
	//router.DELETE("api/users/:id", deleteUser)

	router.Run()
}

func getUsers(ctx *gin.Context) {
	users, _ := fetchUsers()
	ctx.IndentedJSON(200, users)
}

func fetchUser(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Param("id"))
	user, _ := fetchUserById(id)
	ctx.IndentedJSON(200, user)
}

func saveUser(ctx *gin.Context) {
	var newUser user
	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	id, _ := persistUser(newUser)
	ctx.IndentedJSON(200, id)
}

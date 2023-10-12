package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	var users []user
	dsn := "root:localhost@tcp(127.0.0.1:3306)/mind_sharer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("%v", err)
	}
	result := db.Find(&users)
	fmt.Println(result.RowsAffected) // returns found records count, equals `len(users)`
	fmt.Println(result.Error)        // returns error
	ctx.IndentedJSON(200, users)
}

func fetchUser(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Param("id"))
	user, _ := fetchUserById(id)
	ctx.IndentedJSON(200, user)
}

func saveUser(ctx *gin.Context) {
	var newUser user
	fmt.Println(newUser)
	if err := ctx.BindJSON(&newUser); err != nil {
		fmt.Println("%v", err)
		return
	}
	fmt.Println(newUser)

	persistedUser, _ := persistUser(newUser)
	ctx.IndentedJSON(200, persistedUser)
}

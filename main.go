package main

func main() {
	configureMySQLDatabase()
	configureUsersRouter()
}

//func fetchUsers(context *gin.Context) {
//
//	context.IndentedJSON(http.StatusOK, users)
//}

//func fetchUserById(context *gin.Context) {
//	println("start")
//	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
//	for _, currentUser := range users {
//		if currentUser.Id == int(id) {
//			fmt.Println("%s", currentUser)
//			context.IndentedJSON(http.StatusOK, currentUser)
//			return
//		}
//	}
//	println("end")
//}
//
//func saveUser(context *gin.Context) {
//	var newUser user
//
//	// Call BindJSON to bind the received JSON to
//	// newAlbum.
//	if err := context.BindJSON(&newUser); err != nil {
//		return
//	}
//
//	// Add the new album to the slice.
//	users = append(users, newUser)
//	context.IndentedJSON(http.StatusCreated, newUser)
//}
//
//func updateUser(context *gin.Context) {
//	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
//	var userToUpdate user
//	if err := context.BindJSON(&userToUpdate); err != nil {
//		return
//	}
//
//	for i := 0; i < len(users); i++ {
//		if users[i].Id == int(id) {
//			users[i] = userToUpdate
//			return
//		}
//	}
//}
//
//func deleteUser(context *gin.Context) {
//	id := uuid.MustParse(context.Param("id"))
//	for index, currentUser := range users {
//		if currentUser.Id == id {
//			users = append(users[:index], users[index+1:]...)
//		}
//	}
//}

//var users = []user{
//	{
//		Id:       uuid.UUID{},
//		Email:    "",
//		Username: "",
//		Password: "",
//		Posts:    nil,
//	},
//	{
//		Id:       uuid.UUID{},
//		Email:    "",
//		Username: "",
//		Password: "",
//		Posts:    nil,
//	},
//	{
//		Id:       uuid.UUID{},
//		Email:    "",
//		Username: "",
//		Password: "",
//		Posts:    nil,
//	},
//}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	//const path = "file::memory:?cache=shared"
	const path = "test.db"
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	defer func() {
		if db != nil {
			db.Commit()
		}
	}()

	if err != nil {
		panic("Failed to connect database!")
	}

	if err := db.AutoMigrate(
		&User{},
		&Task{},
		&Device{},
		&DeviceTaskRelation{},
		&UserRelation{},
		&Notification{}); err != nil {
		println(err.Error())
	}

	engine := gin.Default()

	engine.POST("/user/login", func(context *gin.Context) {

		user := User{}

		if err := context.BindJSON(&user); err != nil {
			bindErr := fmt.Errorf("BindJSON Error %v", err)
			context.JSON(http.StatusOK, Result{
				Success: false,
				Message: bindErr.Error(),
			})
			return
		}

		var users []User
		db.Select("Id").Where("email=? and password=?", user.Email, user.Password).First(&users)
		l := len(users)

		if l == 0 {
			context.JSON(http.StatusOK, Result{
				Success: false,
				Message: fmt.Sprintf("No such user email='%v',password='%v'", user.Email, user.Password),
			})
			return
		}

		if l == 1 {
			context.JSON(http.StatusOK, DataResult[User]{
				Success: true,
				Message: "Success login.",
				Data: User{
					Id:    users[0].Id,
					Token: "Test Token",
				},
			})
			return
		}

		{
			dbErr := fmt.Errorf("more than one user: %v", user)
			context.JSON(http.StatusOK, Result{
				Success: false,
				Message: dbErr.Error(),
			})
			panic(dbErr)
			return
		}

	})

	engine.GET("/test/getTestMessage", func(context *gin.Context) {
		context.JSON(http.StatusOK, Result{
			Success: true,
		})
	})

	engine.GET("/test/getUsers", func(context *gin.Context) {
		context.JSON(http.StatusOK, DataResult[[]User]{
			Success: true,
			Data: []User{
				{Id: 1},
				{Id: 2},
				{Id: 3},
			},
		})
	})

	engine.GET("/test/getUsersDirect", func(context *gin.Context) {
		context.JSON(http.StatusOK, []User{
			{Id: 1},
			{Id: 2},
			{Id: 3},
		})
	})

	if err := engine.Run(); err != nil {
		println(err.Error())
		panic(err.Error())
	}

}

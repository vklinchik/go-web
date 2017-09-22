package routes

import (
	"gopkg.in/kataras/iris.v6"
	"log"
	"xlog"
	"errors"
)

type User struct {
	Id int
	Name string
}

var users = makeUsers()

func makeUsers() []User {
	users := make([]User, 5) 
	users[0] = User{Id: 101, Name: "user 1"}
	users[1] = User{Id: 102, Name: "user 2"}
	users[2] = User{Id: 103, Name: "user 3"}
	users[3] = User{Id: 104, Name: "user 4"}
	users[4] = User{Id: 105, Name: "user 5"}
	return users
}

func getUser(id int) *User {
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			u := users[i]
			return &u
		}
	}
	return nil
}

func nextId() int {
	maxId := 0
	for i := 0; i < len(users); i++ {
		if users[i].Id > maxId {
			maxId = users[i].Id
		}
	}
	maxId++
	return maxId
}

func GetUser() func(*iris.Context) {
	return func(ctx *iris.Context) {
		xlog.Debug("Executed GetUser")
		id, err := ctx.ParamInt("id")
		if err != nil {
			log.Println(err)
		}

		userp := getUser(id)
		if userp == nil {
			err = errors.New("User not found.")
			log.Println(err)
			ctx.JSON(iris.StatusOK, map[string]string{"error": "user not found"})
		} else {
			ctx.JSON(iris.StatusOK, *userp)
		}
	}
}

func AddUser() func(*iris.Context) {
	return func(ctx *iris.Context) {
		user := User{}
		err := ctx.ReadJSON(&user)
		if err != nil {
			log.Println(err)
			ctx.JSON(iris.StatusBadRequest, map[string]string{"error": "bad request"})
		} else {
			user.Id = nextId()
			users = append(users, user)
			ctx.JSON(iris.StatusOK, user)
		}
	}
}

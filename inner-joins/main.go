package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID                 int
	Username, Password string
}

type Post struct {
	ID   int
	Post string
}

type User_Post struct {
	User_ID, Post_ID int
}

func main() {

	db, err := gorm.Open(sqlite.Open("databases.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{}, &Post{}, &User_Post{})

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&User{ID: 1, Username: "cfabrica46", Password: "01234"})

	db.Create(&Post{ID: 1, Post: "uwu"})

	db.Create(&User_Post{User_ID: 1, Post_ID: 1})

	var user User

	db.Table("user_posts").Select("users.id,users.username,users.password").Joins("INNER JOIN users ON users.id=user_posts.user_id").Joins("INNER JOIN posts ON posts.id=user_posts.post_id").Where("posts.id=?", 1).Scan(&user)

	fmt.Println(user)
}

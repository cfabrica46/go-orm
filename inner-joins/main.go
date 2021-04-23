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
	Posts              []Post `gorm:"many2many:user_posts;foreignKey:ID;joinForeignKey:User_ID;References:User_ID;JoinReferences:User_ID"`
}

type Post struct {
	ID, User_ID int
	Post        string
}

func main() {

	user := User{
		ID:       1,
		Username: "cfabrica46",
		Password: "01234",
		Posts: []Post{
			{User_ID: 1, Post: "uwu"},
			{User_ID: 1, Post: "owo"},
		},
	}

	db, err := gorm.Open(sqlite.Open("databases.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{}, &Post{})

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&user)

	var posts []Post

	err = db.Model(&user).Association("Posts").Find(&posts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(posts)
}

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 int
	Username, Password string
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: 1, Username: "cfabrica46", Password: "01234"})

	var user User

	db.First(&user, "username=?", "cfabrica46")

	fmt.Printf("DEBUGG 1: %d: %s :%s\n", user.ID, user.Username, user.Password)

	db.Model(&user).Update("password", "0000")

	fmt.Printf("DEBUGG 2: %d: %s :%s\n", user.ID, user.Username, user.Password)

	db.Delete(&user, 1)

	db.First(&user, "username=?", "cfabrica46")

	fmt.Printf("DEBUGG 2: %d: %s :%s\n", user.ID, user.Username, user.Password)

}

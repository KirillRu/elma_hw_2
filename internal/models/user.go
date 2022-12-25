package models

import "fmt"

type User struct {
	Id        Uuid   `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"`
	Purchases uint   `json:"purchases"` //number of purchases
}

var UserUpdatesCh chan string

func init() {
	UserUpdatesCh = make(chan string)
}

func (user *User) Log(message string) {
	UserUpdatesCh <- fmt.Sprintf("User:%s (%s), message: %s", user.Name, user.Id, message)
}

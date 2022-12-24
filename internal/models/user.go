package models

import "fmt"

type User struct {
	Id        Uuid   `json:"id"`
	Face      string `json:"face"`
	Purchases uint   `json:"purchases"` //number of purchases
}

var UserUpdatesCh chan string

func init() {
	UserUpdatesCh = make(chan string)
}

func (user User) Log(message string) {
	UserUpdatesCh <- fmt.Sprintf("User:%s (%d), message: %s", user.Face, user.Id, message)
}

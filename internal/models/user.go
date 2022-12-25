package models

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type User struct {
	Id        Uuid      `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	Purchases uint      `json:"purchases"` //number of purchases
}

type UserRespone struct {
	User
	AccessToken string `json:"token"`
}

var UserUpdatesCh chan string
var jwtSecretKey = []byte("veryVerySecretKey")

func init() {
	UserUpdatesCh = make(chan string)
}

func (user *User) Log(message string) {
	UserUpdatesCh <- fmt.Sprintf("User:%s (%s), message: %s", user.Name, user.Id, message)
}

func (user *User) FromRequest(r *http.Request) {
	user.Name = r.FormValue("name")
	user.Login = r.FormValue("login")
	user.Password = r.FormValue("password")
	user.Phone = r.FormValue("phone")
	t, err := time.Parse("01/02/2006", r.FormValue("birth_date"))
	if err == nil {
		user.BirthDate = t
	} else {
		user.BirthDate = time.UnixMilli(0)
	}
	user.Purchases = 0
}

func (user *User) GetToken() string {
	payload := jwt.MapClaims{
		"uid": user.Id,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return ""
	}

	return t
}

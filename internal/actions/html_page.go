package actions

import (
	"fmt"
	"net/http"
)

func GetMain() (string, interface{}, error) {
	return "main.html", nil, nil
}

func GetLogin() (string, interface{}, error) {
	return "login.html", nil, nil
}

func GetUser(r *http.Request) (string, interface{}, error) {
	tokenCookie, err := r.Cookie("access_tocken")
	if err == nil {
		fmt.Println(tokenCookie.Name, tokenCookie.Value)
		user, err := GetUserByToken(tokenCookie.Value)
		if err == nil {
			return "reg.html", *user, nil
		}
	}
	return "reg.html", nil, nil
}

package actions

import "net/http"

func GetMain() (string, interface{}, error) {
	return "main.html", nil, nil
}

func GetLogin() (string, interface{}, error) {
	return "login.html", nil, nil
}

func GetUser(r *http.Request) (string, interface{}, error) {
	return "reg.html", nil, nil
}

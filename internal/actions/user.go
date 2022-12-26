package actions

import (
	"elma_hw_2/internal/models"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"net/http"
)

var users = make(map[models.Uuid]*models.User)
var lastUserId models.Uuid

func GetUserById(userId models.Uuid) (*models.User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, errors.New("Client not found")
}

func Reg(face string) models.Uuid {
	lastUserId = lastUserId.NextNumber()
	user := models.User{
		Id:        lastUserId,
		Name:      face,
		Purchases: 0,
	}
	users[lastUserId] = &user
	users[user.Id].Log("Регистрация")
	return user.Id
}

func RegByUser(user *models.User) (models.UserRespone, error) {
	lastUserId = lastUserId.NextNumber()

	user.Id = lastUserId
	users[user.Id] = user
	users[user.Id].Log("Регистрация")

	resp := models.UserRespone{
		User:        *user,
		AccessToken: user.GetToken(),
	}

	return resp, nil
}

func UpdateByUser(userNew *models.User, r *http.Request) (models.UserRespone, error) {
	fmt.Println(userNew)
	tokenCookie, err := r.Cookie("access_tocken")
	if err == nil {
		userOld, err := GetUserByToken(tokenCookie.Value)
		if err == nil {
			userOld.Name = userNew.Name
			userOld.Login = userNew.Login
			userOld.Password = userNew.Password
			userOld.Phone = userNew.Phone
			userOld.BirthDate = userNew.BirthDate
			userOld.Log("Новые данные клиента сохранены")
			resp := models.UserRespone{
				User:        *userOld,
				AccessToken: tokenCookie.Value,
			}
			return resp, nil
		}

	}

	return models.UserRespone{}, nil
}

func GetUsers() map[models.Uuid]*models.User {
	return users
}

func GetUserByToken(tokenString string) (*models.User, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return models.JwtSecretKey, nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return GetUserById(claims.Uid)
	}
	return nil, errors.New("Client not found")

}

func GetUserByLogin(r *http.Request) (models.UserRespone, error) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	for _, user := range users {
		if user.Login == login && user.Password == password {
			resp := models.UserRespone{
				User:        *user,
				AccessToken: user.GetToken(),
			}
			return resp, nil
		}
	}
	return models.UserRespone{}, errors.New("Нету")
}

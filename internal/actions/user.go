package actions

import (
	"elma_hw_2/internal/models"
	"errors"
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

func GetUsers() map[models.Uuid]*models.User {
	return users
}

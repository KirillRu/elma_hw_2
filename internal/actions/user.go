package actions

import (
	"elma_hw_2/internal/models"
	"errors"
)

var users = make(map[models.Uuid]models.User)

var lastUserId models.Uuid = 0

func GetUserById(userId models.Uuid) (models.User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return models.User{}, errors.New("Client not found")
}

func Reg(face string) models.Uuid {
	lastUserId = lastUserId.NextNumber()
	users[lastUserId] = models.User{
		Id:        lastUserId,
		Face:      face,
		Purchases: 0,
	}
	users[lastUserId].Log("Регистрация")
	return lastUserId
}

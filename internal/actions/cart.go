package actions

import (
	"elma_hw_stat/internal/models"
	"time"
)

var Baskets = make(map[models.Uuid]models.Cart)
var lastCartId models.Uuid = 0

func TakeBasket(u models.User) (models.Cart, error) {
	if c, ok := Baskets[u.Id]; ok {
		return c, nil
	}
	lastCartId = lastCartId.NextNumber()
	Baskets[u.Id] = models.Cart{
		Id:         lastCartId,
		UserId:     u.Id,
		Date:       time.Now(),
		ProductIds: make(map[models.Uuid]uint),
	}
	Baskets[u.Id].Log("The customer took the basket")
	return Baskets[u.Id], nil
}

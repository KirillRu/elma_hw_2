package actions

import (
	"elma_hw_2/internal/models"
	"errors"
	"fmt"
	"time"
)

var Orders []models.Order
var lastOrderId models.Uuid = 0

func Buy(u models.User) error {
	if c, ok := Baskets[u.Id]; ok {
		if len(c.ProductIds) < 1 {
			return errors.New("There are no products in the basket")
		}
		lastOrderId = lastOrderId.NextNumber()
		order := models.Order{
			Id:         lastOrderId,
			UserId:     u.Id,
			Date:       time.Now(),
			ProductIds: make(map[models.Uuid]models.Сost),
		}
		for productId, quantity := range c.ProductIds {
			cost, err := Warehouse[productId].Buy(quantity)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				order.ProductIds[productId] = cost
			}
		}
		if len(order.ProductIds) > 0 {
			order.Log("The order is completed")
			Orders = append(Orders, order)
			c.Clear()
			u.Purchases++
			return nil
		}
		return errors.New("Purchase failed")
	}
	return errors.New("You don't have a basket")
}

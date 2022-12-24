package models

import (
	"fmt"
	"time"
)

var CartUpdatesCh chan string

func init() {
	CartUpdatesCh = make(chan string)
}

type Cart struct {
	Id         Uuid          `json:"id"`
	UserId     Uuid          `json:"user_id"`
	ProductIds map[Uuid]uint `json:"product_ids"`
	Date       time.Time     `json:"date"`
}

func (c Cart) IntoCart(p Product, quantity uint) error {
	err := p.IntoCart(quantity)
	if err != nil {
		return err
	}

	if _, ok := c.ProductIds[p.Id]; ok {
		c.ProductIds[p.Id] += quantity
	} else {
		c.ProductIds[p.Id] = quantity
	}
	c.Date = time.Now()
	c.Log(fmt.Sprintf("The customer put the product %s (%d) in the basket", p.Name, p.Id))
	return nil
}

func (c Cart) Clear() {
	c.ProductIds = map[Uuid]uint{}
	c.Log("The basket is cleared")
}

func (c Cart) Log(message string) {
	UserUpdatesCh <- fmt.Sprintf("Cart: %d for the user: %d, message: %s", c.Id, c.UserId, message)
}

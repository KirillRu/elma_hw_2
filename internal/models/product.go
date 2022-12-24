package models

import (
	"errors"
	"fmt"
)

type Product struct {
	Id       Uuid    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	Reserved uint    `json:"reserved"`
}

var ProductUpdatesCh chan string

func init() {
	ProductUpdatesCh = make(chan string)
}

func (p Product) Log(message string) {
	ProductUpdatesCh <- fmt.Sprintf("Product:%s (%d), message: %s", p.Name, p.Id, message)
}

func (p Product) Buy(quantity uint) (Сost, error) {
	if p.Quantity < quantity {
		return Сost{}, errors.New("Not enough product")
	}
	//if p.Reserved < quantity {
	//	return Сost{}, errors.New("Something went wrong")
	//}
	p.Quantity -= quantity
	p.Reserved -= quantity
	//p.Log()
	return Сost{
		Price:    p.Price,
		Quantity: quantity,
	}, nil
}

func (p Product) IntoCart(quantity uint) error {
	if p.Quantity < quantity+p.Reserved {
		return errors.New("Not enough product")
	}
	p.Reserved += quantity
	p.Log("The customer put the product in the basket")
	return nil
}

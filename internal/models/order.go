package models

import (
	"fmt"
	"time"
)

var OrderUpdatesCh chan string

func init() {
	OrderUpdatesCh = make(chan string)
}

type Order struct {
	Id         Uuid          `json:"id"`
	UserId     Uuid          `json:"user_id"`
	ProductIds map[Uuid]Сost `json:"product_ids"`
	Date       time.Time     `json:"date"`
}

func (o *Order) Log(message string) {
	OrderUpdatesCh <- fmt.Sprintf("Order: %s for the user: %s, message: %s", o.Id, o.UserId, message)
}

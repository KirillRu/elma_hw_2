package main

import (
	"context"
	"elma_hw_2/internal/actions"
	"elma_hw_2/internal/models"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	users := []string{
		"Кирилл",
		"Алексей",
		"Анна",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	go MergeData(ctx)

	for _, face := range users {
		go func(face string) {
			buyProcess(face)
		}(face)
	}
	<-ctx.Done()
}

func buyProcess(face string) {
	userId := actions.Reg(face)
	user, err := actions.GetUserById(userId)

	if err != nil {
		return
	}
	cart, err := actions.TakeBasket(user)
	if err != nil {
		return
	}

	var countProducts = rand.Intn(5) + 1

	for countProducts > 0 { // Условие
		p := actions.GetRandomProduct()
		err = cart.IntoCart(p, uint(rand.Intn(10)+1))
		if err != nil {
			//fmt.Println(err.Error())
		}
		countProducts--
	}
	err = actions.Buy(user)
	if err != nil {
		//fmt.Println(err.Error())
	}
}

func MergeData(ctx context.Context) {

	commonUpdatesCh := make(chan string)
	go func() {
		for {
			select {
			case userMessage := <-models.UserUpdatesCh:
				commonUpdatesCh <- userMessage
			case productMessage := <-models.ProductUpdatesCh:
				commonUpdatesCh <- productMessage
			case cartMessage := <-models.CartUpdatesCh:
				commonUpdatesCh <- cartMessage
			case orderMessage := <-models.OrderUpdatesCh:
				commonUpdatesCh <- orderMessage
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			fmt.Println("COMMON UPDATE ", <-commonUpdatesCh)
		}
	}()

}

package main

import (
	"context"
	"elma_hw_stat/internal/actions"
	"elma_hw_stat/internal/models"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	go MergeData(ctx)

	for _, face := range users {
		go buyProcess(face)
	}
	<-ctx.Done()
}

func buyProcess(face string) {
	userId := actions.Reg(face)
	user, err := actions.GetUserById(userId)
	if err != nil {
		fmt.Println(err.Error())
	}
	cart, err := actions.TakeBasket(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	var countProducts = rand.Intn(5) + 1 // Объявление и инициализация

	for countProducts > 0 { // Условие
		p := actions.GetRandomProduct()
		err = cart.IntoCart(p, uint(rand.Intn(10)+1))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("\n\n%+v\n\n", p)
		countProducts-- // Обратный отсчет; в противном случае цикл будет длиться вечно
	}
	err = actions.Buy(user)
	if err != nil {
		fmt.Println(err.Error())
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

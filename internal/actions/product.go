package actions

import (
	"elma_hw_2/internal/models"
	"errors"
	"math/rand"
)

var Warehouse = make(map[models.Uuid]*models.Product)
var lastProductId models.Uuid

func init() {
	inventory()
}

func GetRandomProduct() *models.Product {
	var num = rand.Intn(len(Warehouse))
	for _, p := range Warehouse {
		if num == 0 {
			return p
		}
		num--
	}
	return nil
}

func GetProductById(productId models.Uuid) (*models.Product, error) {
	if p, ok := Warehouse[productId]; ok {
		return p, nil
	}

	return nil, errors.New("Product not found")
}

func inventory() {
	lastProductId = lastProductId.NextNumber()
	Warehouse[lastProductId] = &models.Product{
		Id:       lastProductId,
		Name:     "The handle is blue",
		Price:    30,
		Quantity: 20,
		Reserved: 0,
	}

	lastProductId = lastProductId.NextNumber()
	Warehouse[lastProductId] = &models.Product{
		Id:       lastProductId,
		Name:     "The handle is red",
		Price:    30,
		Quantity: 20,
		Reserved: 0,
	}

	lastProductId = lastProductId.NextNumber()
	Warehouse[lastProductId] = &models.Product{
		Id:       lastProductId,
		Name:     "The handle is yelow",
		Price:    30,
		Quantity: 20,
		Reserved: 0,
	}

	lastProductId = lastProductId.NextNumber()
	Warehouse[lastProductId] = &models.Product{
		Id:       lastProductId,
		Name:     "The handle is white",
		Price:    30,
		Quantity: 20,
		Reserved: 0,
	}

	lastProductId = lastProductId.NextNumber()
	Warehouse[lastProductId] = &models.Product{
		Id:       lastProductId,
		Name:     "The handle is black",
		Price:    30,
		Quantity: 20,
		Reserved: 0,
	}
}

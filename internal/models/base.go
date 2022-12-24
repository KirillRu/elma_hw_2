package models

import "github.com/google/uuid"

type Uuid string

type Сost struct {
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}

func (id Uuid) NextNumber() Uuid {
	return Uuid(uuid.New().String())
}

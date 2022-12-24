package models

type Uuid uint64

type Сost struct {
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}

func (id Uuid) NextNumber() Uuid {
	id++
	return id
}

//type Statistic interface {
//	Log(message string)
//}
//
//func Log(stat Statistic, message string) {
//	stat.Log(message)
//}

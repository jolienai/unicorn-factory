package model

type UnicornToProduceRequest struct {
	Amount int `json:"amount" validate:"required,min=1,max=100000"`
}

package entity

import "errors"

var (
	ErrOrderNotFound   = errors.New("not found order")
	ErrInvalidOrderID  = errors.New("invalid order id")
	ErrInvalidQuantity = errors.New("quantity must be positive")
	ErrInvalidStatus   = errors.New("invalid order status")
)

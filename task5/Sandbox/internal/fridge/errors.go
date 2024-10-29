package fridge

import "errors"

var (
	ErrNoData         = errors.New("no data available")
	ErrDuplicateID    = errors.New("duplicate product ID")
	ErrInvalidProduct = errors.New("invalid product")
)

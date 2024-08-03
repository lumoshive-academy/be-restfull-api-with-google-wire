package validation

import (
	"errors"
	"restfullapi/collections"
)

func ValidateProduct(product *collections.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	return nil
}

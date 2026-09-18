package service

import "errors"

var ErrProductUnavailable = errors.New("product is unavailable")

type OrderService struct {
	productAvailability ProductAvailability
}

func (o *OrderService) CreateOrder(product string) error {
	available, err := o.productAvailability.IsAvailable(product)
	if err != nil {
		return err
	}
	if !available {
		return ErrProductUnavailable
	}
	return nil
}

type ProductAvailability interface {
	IsAvailable(product string) (bool, error)
}

func NewOrderService(productAvailability ProductAvailability) *OrderService {
	return &OrderService{
		productAvailability: productAvailability,
	}
}

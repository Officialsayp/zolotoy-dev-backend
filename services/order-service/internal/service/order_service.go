package service

import (
	"errors"
	"fmt"

	"github.com/Officialsayp/zolotoy-dev-backend/services/order-service/internal/domain"
	"github.com/google/uuid"
)

type OrderService struct {
	productAvailability ProductAvailability
}

var (
	ErrProductUnavailable = errors.New("product is unavailable")
	ErrInvalidOrder       = errors.New("invalid order")
)

func (o *OrderService) CreateOrder(input CreateOrderInput) (domain.Order, error) {
	items, err := input.toDomainItems()
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w: %v", ErrInvalidOrder, err)
	}

	paymentMethod, err := domain.ParsePaymentMethod(input.PaymentMethod)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w: %v", ErrInvalidOrder, err)
	}

	for _, item := range items {
		available, err := o.productAvailability.IsAvailable(
			item.ProductID,
			item.Quantity,
		)
		if err != nil {
			return domain.Order{}, err
		}

		if !available {
			return domain.Order{}, ErrProductUnavailable
		}
	}

	order, err := domain.NewOrder(
		domain.ID(uuid.NewString()),
		items,
		domain.BuyerID(input.BuyerID),
		paymentMethod,
		domain.DeliveryAddress(input.DeliveryAddress),
		domain.BuyerComment(input.BuyerComment),
	)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w: %v", ErrInvalidOrder, err)
	}

	return order, nil
}

type ProductAvailability interface {
	IsAvailable(productID string, quantity int64) (bool, error)
}

func NewOrderService(productAvailability ProductAvailability) *OrderService {
	return &OrderService{
		productAvailability: productAvailability,
	}
}

type CreateOrderItemInput struct {
	ProductID string
	Name      string
	Quantity  int64
	UnitPrice int64
	Currency  string
}

type CreateOrderInput struct {
	BuyerID         string
	PaymentMethod   string
	Items           []CreateOrderItemInput
	DeliveryAddress string
	BuyerComment    string
}

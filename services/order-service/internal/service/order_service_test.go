package service

import (
	"errors"
	"testing"
)

type fakeProductAvailability struct {
	available bool
	err       error
}

func (f *fakeProductAvailability) IsAvailable(
	_ string,
	_ int64,
) (bool, error) {
	return f.available, f.err
}

func TestOrderService_CreateOrder_AvailableProduct(t *testing.T) {
	availability := &fakeProductAvailability{
		available: true,
	}

	orderService := NewOrderService(availability)

	input := CreateOrderInput{
		BuyerID:       "buyer-123",
		PaymentMethod: "prepaid",
		Items: []CreateOrderItemInput{
			{
				ProductID: "keyboard",
				Name:      "Keyboard",
				Quantity:  1,
				UnitPrice: 10000,
				Currency:  "RUB",
			},
		},
		DeliveryAddress: "Krasnodar",
	}

	_, err := orderService.CreateOrder(input)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestOrderService_CreateOrder_UnavailableProduct(t *testing.T) {
	availability := &fakeProductAvailability{
		available: false,
	}
	orderService := NewOrderService(availability)
	input := CreateOrderInput{
		BuyerID:       "buyer-123",
		PaymentMethod: "prepaid",
		Items: []CreateOrderItemInput{
			{
				ProductID: "keyboard",
				Name:      "Keyboard",
				Quantity:  1,
				UnitPrice: 10000,
				Currency:  "RUB",
			},
		},
		DeliveryAddress: "Krasnodar",
	}

	_, err := orderService.CreateOrder(input)
	if !errors.Is(err, ErrProductUnavailable) {
		t.Fatalf("expected ErrProductUnavailable, got %v",
			err,
		)
	}
}
func TestOrderService_CreateOrder_AvailabilityError(t *testing.T) {
	expectedErr := errors.New("availability check failed")
	availability := &fakeProductAvailability{
		err: expectedErr,
	}
	orderService := NewOrderService(availability)
	input := CreateOrderInput{
		BuyerID:       "buyer-123",
		PaymentMethod: "prepaid",
		Items: []CreateOrderItemInput{
			{
				ProductID: "keyboard",
				Name:      "Keyboard",
				Quantity:  1,
				UnitPrice: 10000,
				Currency:  "RUB",
			},
		},
		DeliveryAddress: "Krasnodar",
	}

	_, err := orderService.CreateOrder(input)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected availability error, got %v", err)
	}
}

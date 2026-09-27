package service

import (
	"errors"
	"testing"
)

type fakeProductAvailability struct {
	available bool
	err       error
}

func (f *fakeProductAvailability) IsAvailable(product string) (bool, error) {
	return f.available, f.err
}

func TestOrderService_CreateOrder_AvailableProduct(t *testing.T) {
	availability := &fakeProductAvailability{
		available: true,
	}
	orderService := NewOrderService(availability)
	err := orderService.CreateOrder("keyboard")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestOrderService_CreateOrder_UnavailableProduct(t *testing.T) {
	availability := &fakeProductAvailability{
		available: false,
	}
	orderService := NewOrderService(availability)
	err := orderService.CreateOrder("keyboard")
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
	err := orderService.CreateOrder("keyboard")
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected availability error, got %v", err)
	}
}

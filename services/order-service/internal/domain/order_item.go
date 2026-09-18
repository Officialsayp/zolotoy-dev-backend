package domain

import "errors"

type OrderItem struct {
	SKU              SKU
	Quantity         int64
	UnitPrice        Money
	Discount         Money
	ReservedQuantity int64
}

func NewOrderItem(
	sku SKU,
	quantity int64,
	unitPrice Money,
	discount Money,
) (OrderItem, error) {
	switch {
	case sku == "":
		return OrderItem{}, errors.New("sku is empty")
	case quantity <= 0:
		return OrderItem{}, errors.New("quantity is empty")
	case unitPrice.Amount < 0:
		return OrderItem{}, errors.New("incorrect value unit price")
	case discount.Amount < 0:
		return OrderItem{}, errors.New("incorrect value discount")
	case discount.Amount > unitPrice.Amount:
		return OrderItem{}, errors.New("incorrect value discount")
	}
	return OrderItem{
		SKU:              sku,
		Quantity:         quantity,
		UnitPrice:        unitPrice,
		Discount:         discount,
		ReservedQuantity: 0,
	}, nil
}

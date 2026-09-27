package domain

import "errors"

type OrderItem struct {
	ProductID           string
	ProductNameSnapshot string
	Quantity            int64
	UnitPrice           Money
	TotalPrice          Money
}

func NewOrderItem(
	productID string,
	productName string,
	quantity int64,
	unitPrice Money,
) (OrderItem, error) {
	switch {
	case productID == "":
		return OrderItem{}, errors.New("product id is empty")

	case productName == "":
		return OrderItem{}, errors.New("product name is empty")

	case quantity <= 0:
		return OrderItem{}, errors.New(
			"quantity must be greater than zero",
		)

	case unitPrice.Amount <= 0:
		return OrderItem{}, errors.New(
			"unit price must be greater than zero",
		)
	}

	totalPrice, err := NewMoney(
		unitPrice.Amount*quantity,
		unitPrice.Currency,
	)
	if err != nil {
		return OrderItem{}, err
	}

	return OrderItem{
		ProductID:           productID,
		ProductNameSnapshot: productName,
		Quantity:            quantity,
		UnitPrice:           unitPrice,
		TotalPrice:          totalPrice,
	}, nil
}

package service

import "github.com/Officialsayp/zolotoy-dev-backend/services/order-service/internal/domain"

func (input CreateOrderInput) toDomainItems() ([]domain.OrderItem, error) {
	items := make([]domain.OrderItem, 0, len(input.Items))

	for _, item := range input.Items {
		unitPrice, err := domain.NewMoney(
			item.UnitPrice,
			domain.Currency(item.Currency),
		)
		if err != nil {
			return nil, err
		}

		orderItem, err := domain.NewOrderItem(
			item.ProductID,
			item.Name,
			item.Quantity,
			unitPrice,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, orderItem)
	}

	return items, nil
}

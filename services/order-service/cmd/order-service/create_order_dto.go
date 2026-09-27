package main

import (
	"time"

	"github.com/Officialsayp/zolotoy-dev-backend/services/order-service/internal/domain"
	"github.com/Officialsayp/zolotoy-dev-backend/services/order-service/internal/service"
)

type createOrderItemRequest struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int64  `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
	Currency  string `json:"currency"`
}
type createOrderV1Request struct {
	BuyerID         string                   `json:"buyer_id"`
	PaymentMethod   string                   `json:"payment_method"`
	Items           []createOrderItemRequest `json:"items"`
	DeliveryAddress string                   `json:"delivery_address"`
	BuyerComment    string                   `json:"buyer_comment"`
}

func (r createOrderV1Request) toServiceInput() service.CreateOrderInput {
	items := make([]service.CreateOrderItemInput, 0, len(r.Items))

	for _, item := range r.Items {
		items = append(items, service.CreateOrderItemInput{
			ProductID: item.ProductID,
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			Currency:  item.Currency,
		})
	}

	return service.CreateOrderInput{
		BuyerID:         r.BuyerID,
		PaymentMethod:   r.PaymentMethod,
		Items:           items,
		DeliveryAddress: r.DeliveryAddress,
		BuyerComment:    r.BuyerComment,
	}
}

type createOrderItemResponse struct {
	ProductID  string `json:"product_id"`
	Name       string `json:"name"`
	Quantity   int64  `json:"quantity"`
	UnitPrice  int64  `json:"unit_price"`
	TotalPrice int64  `json:"total_price"`
	Currency   string `json:"currency"`
}

type createOrderResponse struct {
	ID              string                    `json:"id"`
	BuyerID         string                    `json:"buyer_id"`
	Status          string                    `json:"status"`
	PaymentStatus   string                    `json:"payment_status"`
	PaymentMethod   string                    `json:"payment_method"`
	Items           []createOrderItemResponse `json:"items"`
	DeliveryAddress string                    `json:"delivery_address"`
	BuyerComment    string                    `json:"buyer_comment"`
	CreatedAt       time.Time                 `json:"created_at"`
}

func newCreateOrderResponse(order domain.Order) createOrderResponse {
	domainItems := order.Items()

	items := make([]createOrderItemResponse, 0, len(domainItems))

	for _, item := range domainItems {
		items = append(items, createOrderItemResponse{
			ProductID:  item.ProductID,
			Name:       item.ProductNameSnapshot,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice.Amount,
			TotalPrice: item.TotalPrice.Amount,
			Currency:   string(item.UnitPrice.Currency),
		})
	}

	return createOrderResponse{
		ID:              string(order.ID()),
		BuyerID:         string(order.BuyerID()),
		Status:          string(order.Status()),
		PaymentStatus:   string(order.PaymentStatus()),
		PaymentMethod:   string(order.PaymentMethod()),
		Items:           items,
		DeliveryAddress: string(order.DeliveryAddress()),
		BuyerComment:    string(order.BuyerComment()),
		CreatedAt:       order.CreatedAt(),
	}
}

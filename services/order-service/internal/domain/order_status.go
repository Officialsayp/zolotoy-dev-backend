package domain

type OrderStatus string

const (
	OrderStatusCreated      OrderStatus = "created"
	OrderStatusProcessing   OrderStatus = "processing"
	OrderStatusShipped      OrderStatus = "shipped"
	OrderStatusDelivered    OrderStatus = "delivered"
	OrderStatusCompleted    OrderStatus = "completed"
	OrderStatusCancellation OrderStatus = "cancellation"
	OrderStatusCancelled    OrderStatus = "cancelled"
)

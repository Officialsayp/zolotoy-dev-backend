package domain

type OrderPaymentStatus string

const (
	OrderPaymentStatusAwaitingPayment  OrderPaymentStatus = "awaiting_payment"
	OrderPaymentStatusProcessing       OrderPaymentStatus = "processing"
	OrderPaymentStatusPaid             OrderPaymentStatus = "paid"
	OrderPaymentStatusRefundProcessing OrderPaymentStatus = "refund_processing"
	OrderPaymentStatusRefunded         OrderPaymentStatus = "refunded"
	OrderPaymentStatusFailed           OrderPaymentStatus = "failed"
	OrderPaymentStatusCancelled        OrderPaymentStatus = "cancelled"
)

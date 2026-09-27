package domain

import "errors"

type PaymentMethod string

const (
	PaymentMethodPrepaid            PaymentMethod = "prepaid"
	PaymentMethodPayOnReceiptOnline PaymentMethod = "pay_on_receipt_online"
)

func ParsePaymentMethod(value string) (PaymentMethod, error) {
	method := PaymentMethod(value)

	switch method {
	case PaymentMethodPrepaid,
		PaymentMethodPayOnReceiptOnline:
		return method, nil
	default:
		return "", errors.New("unsupported payment method")
	}
}

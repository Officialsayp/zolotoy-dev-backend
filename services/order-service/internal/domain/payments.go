package domain

type PaymentMethod string

const (
	PaymentMethodDebitCard  PaymentMethod = "debit_card"
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodSBP        PaymentMethod = "sbp"
)

type PaymentTiming string

const (
	PaymentTimingPrepaid   PaymentTiming = "prepaid"
	PaymentTimingOnReceipt PaymentTiming = "on_receipt"
)

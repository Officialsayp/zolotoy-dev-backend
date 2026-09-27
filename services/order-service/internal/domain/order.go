package domain

import (
	"errors"
	"time"
)

type Order struct {
	id              ID
	createdAt       time.Time
	items           []OrderItem
	status          OrderStatus
	paymentStatus   OrderPaymentStatus
	buyerID         BuyerID
	paymentMethod   PaymentMethod
	deliveryAddress DeliveryAddress
	deliveryCost    Money
	buyerComment    BuyerComment
}

func NewOrder(id ID, items []OrderItem, buyerID BuyerID, paymentMethod PaymentMethod, deliveryAddress DeliveryAddress, buyerComment BuyerComment) (Order, error) {
	switch {
	case id == "":
		return Order{}, errors.New("id is empty")
	case len(items) == 0:
		return Order{}, errors.New("items is empty")
	case buyerID == "":
		return Order{}, errors.New("buyer id is empty")
	case paymentMethod == "":
		return Order{}, errors.New("payment method is empty")
	case deliveryAddress == "":
		return Order{}, errors.New("delivery address is empty")
	}

	copyItems := make([]OrderItem, len(items))
	copy(copyItems, items)
	deliveryCost, err := NewMoney(0, CurrencyRUB)
	if err != nil {
		return Order{}, err
	}
	return Order{
		id:              id,
		createdAt:       time.Now(),
		items:           copyItems,
		status:          OrderStatusCreated,
		paymentStatus:   OrderPaymentStatusAwaitingPayment,
		buyerID:         buyerID,
		paymentMethod:   paymentMethod,
		deliveryAddress: deliveryAddress,
		deliveryCost:    deliveryCost,
		buyerComment:    buyerComment,
	}, nil
}

func (o *Order) ToPay() error {
	switch {
	case o.status == OrderStatusCancelled:
		return errors.New("cancelled order cannot be paid")
	case o.status == OrderStatusCompleted:
		return errors.New("completed order cannot be paid")
	case o.paymentStatus != OrderPaymentStatusAwaitingPayment &&
		o.paymentStatus != OrderPaymentStatusFailed:
		return errors.New("order is not awaiting payment or failed status")
	}

	o.paymentStatus = OrderPaymentStatusPaid

	return nil
}

func (o *Order) RequestCancellation() error {
	switch {
	case o.paymentStatus == OrderPaymentStatusRefunded:
		return errors.New("the order should not be in the refunded status")
	case o.status == OrderStatusCancellation:
		return errors.New("the order should not be in the cancellation status")
	case o.status == OrderStatusCancelled:
		return errors.New("the order should not be in the cancelled status")
	case o.status == OrderStatusCompleted:
		return errors.New("the order should not be in the completed status")
	case o.status == OrderStatusShipped:
		return errors.New("the order should not be in the shipped status")
	}
	o.status = OrderStatusCancellation
	if o.paymentStatus == OrderPaymentStatusPaid {
		o.paymentStatus = OrderPaymentStatusRefundProcessing
	}

	return nil
}

func (o *Order) CompleteRefund() error {
	switch {
	case o.status != OrderStatusCancellation:
		return errors.New("order cancellation is not in progress")
	case o.paymentStatus != OrderPaymentStatusRefundProcessing:
		return errors.New("refund is not in progress")
	}

	o.paymentStatus = OrderPaymentStatusRefunded
	return nil
}

func (o *Order) Cancel() error {
	switch {
	case o.status != OrderStatusCancellation:
		return errors.New("first you need to the order cancellation status")
	case o.paymentStatus != OrderPaymentStatusRefunded:
		return errors.New("first you need to complete refund")
	}
	o.status = OrderStatusCancelled
	o.paymentStatus = OrderPaymentStatusCancelled
	return nil
}

func (o *Order) StartPayment() error {
	switch o.paymentMethod {
	case PaymentMethodPrepaid:
		if o.status != OrderStatusCreated {
			return errors.New(
				"prepaid payment can only be started for a created order",
			)
		}

	case PaymentMethodPayOnReceiptOnline:
		if o.status != OrderStatusDelivered {
			return errors.New(
				"payment on receipt can only be started after delivery",
			)
		}

	default:
		return errors.New("unknown payment method")
	}
	o.paymentStatus = OrderPaymentStatusProcessing
	return nil
}

func (o *Order) StartProcessing() error {
	switch o.paymentMethod {
	case PaymentMethodPrepaid:
		if o.paymentStatus != OrderPaymentStatusPaid {
			return errors.New(
				"prepaid order must be paid before processing",
			)
		}

	case PaymentMethodPayOnReceiptOnline:
		if o.paymentStatus != OrderPaymentStatusAwaitingPayment {
			return errors.New(
				"payment-on-receipt order must be awaiting payment before processing",
			)
		}

	default:
		return errors.New("unknown payment method")
	}
	o.status = OrderStatusProcessing
	return nil
}

func (o *Order) Ship() error {
	if o.status != OrderStatusProcessing {
		return errors.New("order is not in processing status")
	}

	o.status = OrderStatusShipped
	return nil
}

func (o *Order) Deliver() error {
	if o.status != OrderStatusShipped {
		return errors.New("order is not in shipped status")
	}

	o.status = OrderStatusDelivered
	return nil
}

func (o *Order) Complete() error {
	switch {
	case o.status != OrderStatusDelivered:
		return errors.New("order is cannot delivered")
	case o.paymentStatus != OrderPaymentStatusPaid:
		return errors.New("order is cannot paid")
	}
	o.status = OrderStatusCompleted
	return nil
}

func (o *Order) Failed() error {
	if o.paymentStatus == OrderPaymentStatusProcessing || o.paymentStatus == OrderPaymentStatusRefundProcessing {
		o.paymentStatus = OrderPaymentStatusFailed
	} else {
		return errors.New("the payment status must be processing or refund processing")
	}
	return nil
}

func (o Order) ID() ID {
	return o.id
}

func (o Order) BuyerID() BuyerID {
	return o.buyerID
}

func (o Order) Status() OrderStatus {
	return o.status
}

func (o Order) PaymentStatus() OrderPaymentStatus {
	return o.paymentStatus
}

func (o Order) PaymentMethod() PaymentMethod {
	return o.paymentMethod
}

func (o Order) Items() []OrderItem {
	items := make([]OrderItem, len(o.items))
	copy(items, o.items)

	return items
}

func (o Order) DeliveryAddress() DeliveryAddress {
	return o.deliveryAddress
}

func (o Order) BuyerComment() BuyerComment {
	return o.buyerComment
}

func (o Order) CreatedAt() time.Time {
	return o.createdAt
}

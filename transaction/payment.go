package transaction

import "time"

type Payment struct {
	PaymentID string
	Method    string
	PaidAt    time.Time
	CancelAt  time.Time
	ExpiredAt time.Time
}

package types

type Money int64
type PAN string
type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Card struct {
	id int
	balance Money
	pan PAN
	currency Currency
}
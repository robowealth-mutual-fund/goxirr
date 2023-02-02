/*
Package goxirr is a simple implementation of a function for calculating
the Internal Rate of Return for irregular cash flow (XIRR).
*/
package goxirr

import (
	"math"
	"time"
)

// A Transaction represent a single transaction from a series of irregular payments.
type Transaction struct {
	Date time.Time
	Cash float64
}

type Options func(*Option)

type Option struct {
	round bool
	digit float64
}

// WithRound options for calculate with round irr
func WithRound(decimalPoint int) Options {
	return func(o *Option) {
		o.round = true
		o.digit = math.Pow(10, float64(decimalPoint))
	}
}

// Transactions represent a cash flow consisting of individual transactions
type Transactions []Transaction

// Xirr returns the Internal Rate of Return (IRR) for an irregular series of cash flows (XIRR)
func Xirr(transactions Transactions, opts ...Options) float64 {
	var o = &Option{}
	for _, opt := range opts {
		opt(o)
	}

	var years []float64
	for _, t := range transactions {
		years = append(years, (t.Date.Sub(transactions[0].Date).Hours()/24)/365)
	}

	residual := 1.0
	step := 0.05
	guess := 0.05
	epsilon := 0.0001
	limit := 10000

	for math.Abs(residual) > epsilon && limit > 0 {
		limit--

		residual = 0.0

		for i, t := range transactions {
			residual += t.Cash / math.Pow(guess, years[i])
		}

		if math.Abs(residual) > epsilon {
			if residual > 0 {
				guess += step
			} else {
				guess -= step
				step /= 2.0
			}
		}
	}

	irr := (guess - 1) * 100
	if o.round {
		return math.Round(irr*o.digit) / o.digit
	}

	return irr
}

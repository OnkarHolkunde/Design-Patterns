package main

import "fmt"

type Payment interface {
	Pay(amount float64) string
}

type Stripe struct{}

func (s *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f via Stripe", amount)
}

type Paypal struct{}

func (p *Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f via Paypal", amount)
}

// Factory function
func ProcessPayment(method string) Payment {
	switch method {
	case "stripe":
		return &Stripe{}

	case "paypal":
		return &Paypal{}

	default:
		return nil
	}
}
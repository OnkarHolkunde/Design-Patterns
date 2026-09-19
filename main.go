package main

import "fmt"

func main() {
	GetDatabaseInstance()
	payment := ProcessPayment("stripe")

	if payment != nil {
		fmt.Println(payment.Pay(100))
	}

}
package main

import (
	"log"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

const (
	customerID = "cus_8AAkt2sqwiZDtZ"
)

var (
	secretKey string
)

func init() {
	secretKey = os.Getenv("STRIPE_SECKEY")
	stripe.Key = secretKey
}

func main() {
	ch, err := charge.New(&stripe.ChargeParams{
		Amount:   2000,
		Currency: "jpy",
		Customer: customerID,
	})
	if err != nil {
		log.Println("Failed to charge a credit card")
		log.Fatal(err)
		return
	}
	log.Println(ch.ID)

	ch, err = charge.New(&stripe.ChargeParams{
		Amount:   1500,
		Currency: "jpy",
		Customer: customerID,
	})
	if err != nil {
		log.Println("Failed to charge a credit card")
		log.Fatal(err)
		return
	}

	log.Println(ch.ID)
}

package main

import (
	"log"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
)

const (
	subscriptionID = "sub_8AFZXwqBhHHcwe"
	customerID     = "cus_8AFZPeoWh6FVZ6"
)

var (
	publishableKey string
	secretKey      string
)

func init() {
	publishableKey = os.Getenv("STRIPE_PUBKEY")
	secretKey = os.Getenv("STRIPE_SECKEY")
	stripe.Key = secretKey
}

func main() {
	s, err := sub.Cancel(subscriptionID, &stripe.SubParams{Customer: customerID})

	if err != nil {
		log.Println("Failed to cancel subscription plan")
		log.Fatal(err)
		return
	}
	log.Println(s.ID)

	it := sub.List(&stripe.SubListParams{Customer: customerID})
	for it.Next() {
		log.Printf("%#v", it.Current())
	}
}

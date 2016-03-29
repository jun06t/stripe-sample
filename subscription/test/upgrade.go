package main

import (
	"log"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
)

var (
	publishableKey string
	secretKey      string
)

const (
	customerID = "cus_8AZfrjEl4ECFmj"
)

func init() {
	publishableKey = os.Getenv("STRIPE_PUBKEY")
	secretKey = os.Getenv("STRIPE_SECKEY")
	stripe.Key = secretKey
}

func main() {
	it := sub.List(&stripe.SubListParams{Customer: customerID})
	if it == nil {
		log.Println("This customer has no subscription")
		return
	}

	var s *stripe.Sub

	for it.Next() {
		s = it.Current().(*stripe.Sub)
	}

	params := &stripe.SubParams{
		Customer: customerID,
		Plan:     "lite",
	}

	result, err := sub.Update(s.ID, params)
	if err != nil {
		log.Println("Failed to update plan")
		log.Fatal(err)
		return
	}
	log.Printf("%#v\n", result)
}

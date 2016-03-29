package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/sub"
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

type TemplateVars struct {
	PublishableKey template.HTML
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./index.html"))
	t.Execute(w, nil)
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	pubKey := TemplateVars{PublishableKey: template.HTML(publishableKey)}
	t := template.Must(template.ParseFiles("./checkout.html"))
	t.Execute(w, pubKey)
}

func chargeHandler(w http.ResponseWriter, r *http.Request) {
	token := r.PostFormValue("stripeToken")
	email := r.PostFormValue("stripeEmail")

	params := &stripe.CustomerParams{
		Email: email,
		Desc:  "Subscription user",
	}
	err := params.SetSource(token)
	if err != nil {
		log.Println("Failed to set token")
		log.Fatal(err)
		return
	}

	customer, err := customer.New(params)
	if err != nil {
		log.Println("Failed to create a customer")
		log.Fatal(err)
		return
	}

	s, err := sub.New(&stripe.SubParams{
		Customer: customer.ID,
		Plan:     "lite",
	})
	if err != nil {
		log.Println("Failed to create subscription")
		log.Fatal(err)
		return
	}

	log.Println(s.ID)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/charge", chargeHandler)
	http.ListenAndServe(":3000", nil)
}

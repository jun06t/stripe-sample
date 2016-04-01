package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/token"
)

const (
	dupCard = "aSHA7lAztCIn6v6T"
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
	t := r.PostFormValue("stripeToken")
	email := r.PostFormValue("stripeEmail")
	to, err := token.Get(t, nil)
	if err != nil {
		log.Println("Failed to create a customer")
		log.Fatal(err)
		return
	}

	fingerprint := to.Card.Fingerprint
	// keep this fingerprint on your database,
	// and you can detect duplicate card.
	if fingerprint == dupCard {
		log.Println("This card has already been used")
		return
	}

	params := &stripe.CustomerParams{
		Email: email,
		Desc:  "Stripe Developer",
	}
	err = params.SetSource(t)
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

	log.Println(customer.ID)
}

func verifyFingerprint() {

}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/charge", chargeHandler)
	http.ListenAndServe(":3000", nil)
}

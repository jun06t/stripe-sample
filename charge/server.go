package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
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
	log.Println(token)

	params := &stripe.ChargeParams{
		Amount:   2000,
		Currency: "jpy",
	}
	err := params.SetSource(token)
	if err != nil {
		log.Println("Failed to set token")
		log.Fatal(err)
		return
	}

	ch, err := charge.New(params)

	if err != nil {
		log.Println("Failed to charge a credit card")
		log.Fatal(err)
		return
	}

	log.Println(ch.ID)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/charge", chargeHandler)
	http.ListenAndServe(":3000", nil)
}

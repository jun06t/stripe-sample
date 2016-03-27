package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

const (
	publishableKey = ""
	secretKey      = ""
)

func init() {
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
		log.Fatal(err)
		os.Exit(1)
	}

	ch, err := charge.New(params)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("%v\n", ch.ID)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/charge", chargeHandler)
	http.ListenAndServe(":3000", nil)
}

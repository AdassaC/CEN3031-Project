package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/customer"
	"github.com/stripe/stripe-go/v71/invoice"
	"github.com/stripe/stripe-go/v71/paymentmethod"
	"github.com/stripe/stripe-go/v71/sub"
	"github.com/stripe/stripe-go/v71/webhook"
)
 
 
 
 
 
 
 func handleRetrieveUpcomingInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		SubscriptionID string `json:"subscriptionId"`
		CustomerID     string `json:"customerId"`
		NewPriceID     string `json:"newPriceId"`
	}
 
 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}
 
 
	s, err := sub.Get(req.SubscriptionID, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("sub.Get: %v", err)
		return
	}
	params := &stripe.InvoiceParams{
		Customer:     stripe.String(req.CustomerID),
		Subscription: stripe.String(req.SubscriptionID),
		SubscriptionItems: []*stripe.SubscriptionItemsParams{{
			ID:         stripe.String(s.Items.Data[0].ID),
			Deleted:    stripe.Bool(true),
			ClearUsage: stripe.Bool(true),
		}, {
			Price:   stripe.String(os.Getenv(req.NewPriceID)),
			Deleted: stripe.Bool(false),
		}},
	}
	in, err := invoice.GetNext(params)
 
 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("invoice.GetNext: %v", err)
		return
	}
 
 
	writeJSON(w, in)
 }
 
 

 
 func handleRetryInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
 
 
	var req struct {
		CustomerID      string `json:"customerId"`
		PaymentMethodID string `json:"paymentMethodId"`
		InvoiceID       string `json:"invoiceId"`
	}
 
 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}
 
 
	// Attach PaymentMethod
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(req.CustomerID),
	}
	pm, err := paymentmethod.Attach(
		req.PaymentMethodID,
		params,
	)
	if err != nil {
		writeJSON(w, struct {
			Error error `json:"error"`
		}{err})
		return
	}
 
 
	// Update invoice settings default
	customerParams := &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(pm.ID),
		},
	}
	c, err := customer.Update(
		req.CustomerID,
		customerParams,
	)
 
 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("customer.Update: %v %s", err, c.ID)
		return
	}
 
 
	// Retrieve Invoice
	invoiceParams := &stripe.InvoiceParams{}
	invoiceParams.AddExpand("payment_intent")
	in, err := invoice.Get(
		req.InvoiceID,
		invoiceParams,
	)
 
 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("invoice.Get: %v", err)
		return
	}
 
 
	writeJSON(w, in)
 }
 
 
 func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("ioutil.ReadAll: %v", err)
		return
	}
 
 
	event, err := webhook.ConstructEvent(b, r.Header.Get("Stripe-Signature"), os.Getenv("STRIPE_WEBHOOK_SECRET"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("webhook.ConstructEvent: %v", err)
		return
	}
 
 
	if event.Type != "checkout.session.completed" {
		return
	}
 
 
	cust, err := customer.Get(event.GetObjectValue("customer"), nil)
	if err != nil {
		log.Printf("customer.Get: %v", err)
		return
	}
 
 
	if event.GetObjectValue("display_items", "0", "custom") != "" &&
		event.GetObjectValue("display_items", "0", "custom", "name") == "Pasha e-book" {
		log.Printf("🔔 Customer is subscribed and bought an e-book! Send the e-book to %s", cust.Email)
	} else {
		log.Printf("🔔 Customer is subscribed but did not buy an e-book.")
	}
 }
 
 
//  func writeJSON(w http.ResponseWriter, v interface{}) {
// 	var buf bytes.Buffer
// 	if err := json.NewEncoder(&buf).Encode(v); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("json.NewEncoder.Encode: %v", err)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	if _, err := io.Copy(w, &buf); err != nil {
// 		log.Printf("io.Copy: %v", err)
// 		return
// 	}
//  }
 
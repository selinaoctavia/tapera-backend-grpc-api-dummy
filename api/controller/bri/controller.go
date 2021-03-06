package bri

import (
	"github.com/gorilla/mux"
	"tapera.mitraintegrasi/service/bri"
)

// Controller struct
type Controller struct {
	srv bri.Service
}

// NewController func
func NewController(srv bri.Service) *Controller {
	return &Controller{srv: srv}
}

// Route func
func (c *Controller) Route(r *mux.Router) {
	//routes grouping
	s := r.PathPrefix("/dummy/bri").Subrouter()
	s.HandleFunc("/pendaftaranpeserta", c.PendaftaranPeserta).Methods("POST")
	s.HandleFunc("/cancelsubscribe", c.CancelSubscribe).Methods("POST")
	s.HandleFunc("/subscription", c.Subscription).Methods("POST")
	s.HandleFunc("/redemption", c.Redemption).Methods("POST")
	s.HandleFunc("/cancelredemption", c.CancelRedemption).Methods("POST")
	s.HandleFunc("/adviceredemptionconfirmation", c.AdviceRedemptionConfirmation).Methods("POST")
	s.HandleFunc("/test", c.Test).Methods("GET")
}

package bkn

import (
	"github.com/gorilla/mux"
	"tapera.mitraintegrasi/service/bkn"
)

// Controller struct
type Controller struct {
	srv bkn.Service
}

// NewController func
func NewController(srv bkn.Service) *Controller {
	return &Controller{srv: srv}
}

// Route func
func (c *Controller) Route(r *mux.Router) {
	//routes grouping
	s := r.PathPrefix("/dummy/bkn").Subrouter()
	s.HandleFunc("/peserta", c.GetPeserta).Methods("POST")
	s.HandleFunc("/riwayatgolonganpeserta", c.GetRiwayatGolonganPeserta).Methods("POST")
}

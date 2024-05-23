package controllers

import (
	"github.com/thapakornd/fiber-go/app/store"
)

type Handler struct {
	userStore    store.UserStore
	addressStore store.AddressStore
	cartStore    store.CartStore
	orderStore   store.OrderStore
	paymentStore store.PaymentStore
	productStore store.ProductStore
	validator    *Validator
}

func NewHandler(
	us store.UserStore,
	as store.AddressStore,
	cs store.CartStore,
	os store.OrderStore,
	ps store.PaymentStore,
	prs store.ProductStore,
) *Handler {
	v := New()
	return &Handler{
		userStore:    us,
		addressStore: as,
		cartStore:    cs,
		orderStore:   os,
		paymentStore: ps,
		productStore: prs,
		validator:    v,
	}
}

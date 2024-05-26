package controllers

import "github.com/thapakornd/fiber-go/app/store"

type Handler struct {
	userStore     store.UserStore
	addressStore  store.AddressStore
	cartStore     store.CartStore
	categoryStore store.CategoryStore
	orderStore    store.OrderStore
	paymentStore  store.PaymentStore
	productStore  store.ProductStore
	validator     *Validator
}

func NewHandler(
	us store.UserStore,
	as store.AddressStore,
	cs store.CartStore,
	css store.CategoryStore,
	os store.OrderStore,
	ps store.PaymentStore,
	prs store.ProductStore,
) *Handler {
	v := New()
	return &Handler{
		userStore:     us,
		addressStore:  as,
		cartStore:     cs,
		categoryStore: css,
		orderStore:    os,
		paymentStore:  ps,
		productStore:  prs,
		validator:     v,
	}
}

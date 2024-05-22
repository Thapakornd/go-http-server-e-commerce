package controllers

import "github.com/thapakornd/fiber-go/app/store"

type Handler struct {
	userStore store.UserStore
	validator *Validator
}

func NewHandler(us store.UserStore) *Handler {
	v := New()
	return &Handler{
		userStore: us,
		validator: v,
	}
}

package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	playerController := ServiceContainer().InjectPlayerController()

	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}", playerController.GetPlayerScore)

	// Thêm route mới
	r.Post("/add-user", playerController.AddUser) // Định nghĩa route cho POST /add-user
	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

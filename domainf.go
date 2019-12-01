package main

import (
	"Domainf/views"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/servers", views.GetServers)
	router.Get("/servers/history", views.GetHistory)
	http.ListenAndServe(":8546", router)
}

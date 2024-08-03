package router

import (
	"restfullapi/middleware"
	"restfullapi/wire"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()

	h := wire.InitializProductHandler()

	r.Use(middleware.Logger)
	r.Use(middleware.BasicAuth)

	r.Get("/products", h.GetAllProducts)
	r.Get("/products/{id}", h.GetProductByID)
	r.Post("/products", h.CreateProduct)
	r.Put("/products/{id}", h.UpdateProduct)
	r.Delete("/products/{id}", h.DeleteProduct)

	return r
}

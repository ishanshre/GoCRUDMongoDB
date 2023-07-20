package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ishanshre/GoCRUDMongoDB/internals/pkg/handlers"
)

// Returns the handler
func Router(h handlers.Handlers) http.Handler {
	// Create a new router
	mux := chi.NewRouter()

	// use cors middleware
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// use logger middleware
	mux.Use(middleware.Logger)

	// defining api path
	mux.Get("/products", h.GetProducts)
	mux.Post("/products", h.CreateProduct)
	mux.Get("/products/{id}", h.GetProduct)
	mux.Delete("/products/{id}", h.DeleteProduct)
	mux.Put("/products/{id}", h.UpdateProduct)

	return mux
}

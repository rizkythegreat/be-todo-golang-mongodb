package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Response struct {
	Id        string    `json:"id" bson:"_id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Msg       string
	Code      int
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // atau "*" untuk semua
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/api", func(router chi.Router) {

		// version 1
		router.Route("/v1", func(router chi.Router) {

			router.Get("/healthcheck", healtCheck)
			router.Get("/todos", getTodos)
			router.Get("/todos/{id}", getTodoById)
			router.Post("/todos/create", createTodo)
			router.Put("/todos/{id}", updateTodo)
			router.Delete("/todos/{id}", deleteTodo)
		})

		// version 2 - add it if you want
		// router.Route("/v2", func(router chi.Router) {
		// })

	})

	return router
}

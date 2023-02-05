package user

import "github.com/go-chi/chi/v5"

func AddUserRoutes(r *chi.Mux, controller *UserController) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", controller.create)
		r.Get("/", controller.getAll)
	})
}

package routes

import (
	"went-template/app/controllers"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func UserRoutes(r *chi.Mux, db *gorm.DB) {

	uc := &controllers.UserController{DB: db}

	r.Get("/users", uc.GetAllUsers)
	r.Get("/users/{id}", uc.GetUserByID)
	r.Post("/users", uc.CreateUser)
	r.Put("/users/{id}", uc.UpdateUser)
	r.Delete("/users/{id}", uc.DeleteUser)
}

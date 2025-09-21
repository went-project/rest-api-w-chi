package routes

import (
	"net/http"
	"went-template/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/gorm"
)

func SetupRoutes(r *chi.Mux, db *gorm.DB) {

	// Basic Health Check Endpoint
	r.Get("/ping", HealthCheck)

	// Swagger setup
	docs.SwaggerInfo.BasePath = "/"
	r.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// User routes
	UserRoutes(r, db)
	// [*RP*] Please do not delete this comment. It is used to automatically add new route files.
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"status": "pong",
	})
}

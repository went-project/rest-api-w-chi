package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"went-template/app/models"
	"went-template/internal/responses"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// GetAllUsers retrieves all users from the database.
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} responses.ErrorResponse
// @Router /users [get]
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}
	render.JSON(w, r, users)
}

// GetUserByID retrieves a user by their ID.
// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} responses.ErrorResponse
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := chi.URLParam(r, "id")
	if id == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, responses.ErrorResponse{Error: "ID parameter is required"})
		return
	}
	if err := uc.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, responses.ErrorResponse{Error: "User not found"})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		}
		return
	}
	render.JSON(w, r, user)
}

// CreateUser creates a new user in the database.
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User to create"
// @Success 201 {object} models.User
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /users [post]
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := user.Validate(); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, user)
}

// UpdateUser updates an existing user in the database.
// @Summary Update an existing user
// @Description Update an existing user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User to update"
// @Success 200 {object} models.User
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := chi.URLParam(r, "id")

	if err := uc.DB.First(&user, id).Error; err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, responses.ErrorResponse{Error: "User not found"})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := user.Validate(); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := uc.DB.Save(&user).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}
	render.JSON(w, r, user)
}

// DeleteUser deletes a user from the database.
// @Summary Delete a user
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := chi.URLParam(r, "id")

	// Validate ID parameter
	if id == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, responses.ErrorResponse{Error: "ID parameter is required"})
		return
	}

	if err := uc.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, responses.ErrorResponse{Error: "User not found"})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		}
		return
	}

	if err := uc.DB.Delete(&user).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, responses.ErrorResponse{Error: err.Error()})
		return
	}
	render.JSON(w, r, map[string]string{"message": "User deleted"})
}

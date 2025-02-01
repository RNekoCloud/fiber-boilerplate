package handlers

import (
	"api-service/http"
	"api-service/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RouterUserManagemet defines routes for user management endpoints.
func (h *Handlers) RouterUserManagemet(app *fiber.App) {
	v1 := app.Group("/api/v1/admin")
	v1.Post("/users", h.Middleware.Protected(), h.CreateUser)
	v1.Get("/users", h.Middleware.Protected(), h.FindUsers)
	v1.Get("/users/:id", h.Middleware.Protected(), h.FindUserById)
	v1.Put("/users/:id", h.Middleware.Protected(), h.EditUserById)
	v1.Delete("/users/:id", h.Middleware.Protected(), h.DeleteUserById)
}

// CreateUser handles the creation of a new user.
func (h *Handlers) CreateUser(c *fiber.Ctx) error {
	// Parse the request body into createUser struct
	var createUser http.UserManagement
	if err := c.BodyParser(&createUser); err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error parsing request create user",
			Data:    nil,
		})
	}

	// Call UserManagement service to create user
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(createUser.Password), 14)

	result, err := h.UserManagement.CreateUser(model.User{
		Name:     createUser.Name,
		Username: createUser.Username,
		Password: string(hashBytes),
		Role:     model.ROLE(createUser.Role),
	})

	if err != nil {
		c.Status(404)
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error creating user request",
			Data:    nil,
		})
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "User created successfully",
		Data:    result,
	})

}

// FindUsers handles the retrieval of all users.
func (h *Handlers) FindUsers(c *fiber.Ctx) error {
	// Call UserManagement service to find all users
	result, err := h.UserManagement.FindUsers()

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error get users",
			Data:    nil,
		})
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "Users found successfully",
		Data:    result,
	})
}

// FindUserById handles the retrieval of a user by ID.
func (h *Handlers) FindUserById(c *fiber.Ctx) error {
	// Call UserManagement service to find user by ID
	result, err := h.UserManagement.FindUserById(c.Params("id"))

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "User id not found",
			Data:    nil,
		})
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "User found successfully",
		Data:    result,
	})
}

// EditUserById handles the update of a user by ID.
func (h *Handlers) EditUserById(c *fiber.Ctx) error {
	// Parse the request body into updateUser struct
	var requestUpdate http.UserManagement
	if err := c.BodyParser(&requestUpdate); err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error parsing request update user",
			Data:    nil,
		})
	}

	// Call UserManagement service to edit user by ID
	result, err := h.UserManagement.EditUser(model.User{
		Name:     requestUpdate.Name,
		Username: requestUpdate.Username,
		Password: requestUpdate.Password,
		Role:     model.ROLE(requestUpdate.Role),
	}, c.Params("id"))

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error updating user request",
			Data:    nil,
		})
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "User updated successfully",
		Data:    result,
	})

}

// DeleteUserById handles the deletion of a user by ID.
func (h *Handlers) DeleteUserById(c *fiber.Ctx) error {
	// Call UserManagement service to delete user by ID
	result, err := h.UserManagement.DeleteUserById(c.Params("id"))

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error deleting user request",
			Data:    nil,
		})
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "User deleted successfully",
		Data:    result,
	})
}

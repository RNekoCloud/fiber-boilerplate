package handlers

import (
	"os"
	"time"

	"api-service/http"
	"api-service/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) RouteAuth(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Post("/register", h.RegisterHandler)
	v1.Post("/login", h.LoginHandler)
	v1.Get("/verify", h.Middleware.Protected(), h.Verify)
}

func (h *Handlers) LoginHandler(c *fiber.Ctx) error {

	var requestLogin http.Login

	if err := c.BodyParser(&requestLogin); err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error parsing request login",
			Data:    nil,
		})
	}

	response, err := h.UserRepository.FindUser(map[string]interface{}{
		"username": requestLogin.Username,
	})

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "User is not registered!",
			Data:    nil,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(requestLogin.Password))

	if err != nil {
		return c.Status(400).JSON(&http.WebResponse{
			Status:  "error",
			Message: "Password is incorrect!",
			Data:    nil,
		})
	}

	// create jwt token from the request username
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": response.Username,
		"user_id":  response.ID,
		"role":     response.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString(h.JwtSecret)

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error generate jwt token",
			Data:    nil,
		})
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "Credentials is valid and Login successfuly!",
		Data: map[string]interface{}{
			"token": token,
		},
	})

}

func (h *Handlers) RegisterHandler(c *fiber.Ctx) error {

	var requestRegister http.Register

	if err := c.BodyParser(&requestRegister); err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error parsing request register",
			Data:    nil,
		})
	}

	response, err := h.UserRepository.CreateUser(model.User{
		Username: requestRegister.Username,
		Password: requestRegister.Password,
		Role:     model.ROLE(requestRegister.Role),
	})

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error registering request",
			Data:    nil,
		})
	}

	// create jwt token from the request username
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": response.Username,
		"user_id":  response.ID,
		"role":     response.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_TOKEN")))

	if err != nil {
		return c.JSON(&http.WebResponse{
			Status:  "error",
			Message: "Error generate jwt token",
			Data:    nil,
		})
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "Register successfully",
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

func (h *Handlers) Verify(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims == nil {
		return c.Status(401).JSON(&http.WebResponse{
			Status:  "error",
			Message: "Unauthorized",
			Data:    nil,
		})
	}

	return c.Status(200).JSON(&http.WebResponse{
		Status:  "success",
		Message: "User verified",
		Data:    claims,
	})

}

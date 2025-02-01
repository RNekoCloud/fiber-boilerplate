package handlers

import (
	"api-service/middleware"
	"api-service/repository"
)

type Handlers struct {
	CourseRepository repository.CourseRepository
	UserRepository   repository.UserRepository
	JwtSecret      []byte
	Middleware       middleware.Middleware
	UserManagement   repository.UserManagementRepository
}

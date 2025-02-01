package repository

import "api-service/model"

type UserManagementRepository interface {
	CreateUser(request model.User) (*model.User, error)
	EditUser(request model.User, id string) (*model.User, error)
	FindUsers() (*[]model.User, error)
	FindUserById(id string) (*model.User, error)
	DeleteUserById(id string) (*model.User, error)
}

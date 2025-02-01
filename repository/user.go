package repository

import (
	"api-service/model"
)

type UserRepository interface {
	CreateUser(request model.User) (*model.User, error)
	FindUser(cond map[string]interface{}) (*model.User, error)
}

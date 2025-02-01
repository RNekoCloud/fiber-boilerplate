package repository

import (
	"api-service/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) UserRepository {
	return &userImpl{
		DB: db,
	}
}

// FindUser implements AuthRepository.
// Use map[string]interface{} for flexbility query.
func (a *userImpl) FindUser(cond map[string]interface{}) (*model.User, error) {
	// check whether the username and password in the database are correct
	var user model.User
	err := a.DB.Where(cond).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// Create User implements AuthRepository.
func (a *userImpl) CreateUser(request model.User) (*model.User, error) {
	// check if the user is already registered
	var user model.User
	err := a.DB.Where("username = ?", request.Username).First(&user).Error
	if err == nil {
		return nil, err
	}

	// hashing password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// create new user
	newUser := model.User{
		Username: request.Username,
		Password: string(hashPassword),
		Role:     model.ROLE(request.Role),
	}

	err = a.DB.Create(&newUser).Error

	if err != nil {
		return nil, err
	}

	// response
	return &newUser, nil

}

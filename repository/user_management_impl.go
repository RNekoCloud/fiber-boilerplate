package repository

import (
	"api-service/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// userManagementImpl is the implementation of UserManagementRepository.
type userManagementImpl struct {
	DB *gorm.DB
}

// NewUserManagementRepository creates a new instance of userManagementImpl.
func NewUserManagementRepository(db *gorm.DB) UserManagementRepository {
	return &userManagementImpl{
		DB: db,
	}
}

// CreateUser creates a new user in the database.
func (u *userManagementImpl) CreateUser(request model.User) (*model.User, error) {
	// Check if the user already exists
	var user model.User
	err := u.DB.Where("username = ?", request.Username).First(&user).Error
	if err == nil {
		return nil, err
	}

	// Generate hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create a new user
	newUser := model.User{
		Name:     request.Name,
		Username: request.Username,
		Password: string(hashPassword),
		Role:     model.ROLE(request.Role),
	}

	err = u.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

// DeleteUserById deletes a user from the database based on the ID.
func (u *userManagementImpl) DeleteUserById(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = u.DB.Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// EditUser edits user information based on the ID.
func (u *userManagementImpl) EditUser(request model.User, id string) (*model.User, error) {
	var hashPassword string
	// Check if there is a request to change the password
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		hashPassword = string(hashedPassword)
	}

	updateUser := model.User{
		Name:     request.Name,
		Username: request.Username,
		Password: hashPassword,
		Role:     model.ROLE(request.Role),
	}

	err := u.DB.Model(&model.User{}).Where("id = ?", id).Updates(&updateUser).Error
	if err != nil {
		return nil, err
	}

	return &updateUser, nil
}

// FindUser returns a list of all users from the database.
func (u *userManagementImpl) FindUsers() (*[]model.User, error) {
	var users []model.User
	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// FindUserById searches for a user by ID.
func (u *userManagementImpl) FindUserById(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

package http

import (
	"time"

	"api-service/helper"
)

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CourseHTTP struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Teacher      *string   `json:"teacher,omitempty"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ThumbnailImg string    `json:"thumbnail_img"`
	PublishedAt  time.Time `json:"published_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CoursePaginationResponse struct {
	Entries    []CourseHTTP       `json:"entries"`
	Pagination *helper.Pagination `json:"pagination"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// User Management Request
type UserManagement struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type EditUserManagementRequest struct {
	Id string `json:"id"`
}

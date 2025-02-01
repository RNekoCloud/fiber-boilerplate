package repository

import (
	"api-service/helper"
	"api-service/model"
)

/*
* This folder handler logic in persistance layer.
* Simply, this is where our query database is doing their job.

* Tips: Param in the method SHOULD NOT BE POINTER.
* Because it has potential error 'POINTER DEFERENCE' when its value is nil.
 */
type CourseRepository interface {
	// Course
	CreateCourse(data model.Course) (*model.Course, error)
	FindCourses(page int, limit int) (*helper.Pagination, []model.Course)
	FindCourse(cond map[string]interface{}) (*model.Course, error)
	DeleteCourse(cond map[string]interface{}) (*model.Course, error)
	EditCourse(cond map[string]interface{}, data model.Course) (*model.Course, error)
}

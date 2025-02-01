package repository

import (
	"api-service/helper"
	"api-service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type courseImpl struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseImpl{
		DB: db,
	}
}

func (repos *courseImpl) CreateCourse(data model.Course) (*model.Course, error) {
	result := repos.DB.Create(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (repos *courseImpl) FindCourses(page int, limit int) (*helper.Pagination, []model.Course) {
	var courses []model.Course
	tx := repos.DB.Model(&model.Course{})

	pagination, txPaginator := helper.Paginator(page, limit, tx)

	txPaginator.Find(&courses)
	if txPaginator.Error != nil {
		logrus.Warnln("[database] There is error:", txPaginator.Error.Error())
		return nil, nil
	}

	return pagination, courses

}

func (repos *courseImpl) FindCourse(cond map[string]interface{}) (*model.Course, error) {
	var course model.Course
	err := repos.DB.Where(cond).First(&course).Error

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (repos *courseImpl) DeleteCourse(cond map[string]interface{}) (*model.Course, error) {
	var course model.Course

	err := repos.DB.Where(cond).First(&course).Error

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (repos *courseImpl) EditCourse(
	cond map[string]interface{},
	data model.Course,
) (*model.Course, error) {
	var course model.Course

	err := repos.DB.Where(cond).First(&course).Error
	if err != nil {
		return nil, err
	}

	data.ID = course.ID
	data.TeacherID = course.TeacherID

	err = repos.DB.Save(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

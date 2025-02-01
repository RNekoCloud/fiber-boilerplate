package handlers

import (
	"strconv"

	"api-service/http"
	"api-service/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

/*
* E Learning API (C) 2024
* This file contains related endpoint routes
* Route should starts with prefix version 'v1' and following with plural nouns.
* Example: /api/v1/courses
 */

/*
* Route setup should naming with pascal case
* Tips: If use word Route before your route name
* Ex: RouteCourses -> Add word 'Route' before `Courses`
 */

func (h *Handlers) RouteCourses(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Post("/courses", h.Middleware.Protected(), h.AddCourse)
	v1.Get("/courses", h.GetCourses)
	v1.Get("/courses/:slug", h.GetDetailCourse)
	v1.Put("/courses/:slug", h.Middleware.Protected(), h.EditCourse)
	v1.Delete("/courses/:slug", h.Middleware.Protected(), h.DeleteCourse)
}

/*
* Handlers name should be explicit and not using buzz word.
* Tips: Use verb for naming function
* Ex: AddCourses(c *fiber.Ctx),
 */

func (h *Handlers) AddCourse(c *fiber.Ctx) error {
	var request http.CourseHTTP
	c.BodyParser(&request)

	teacher := c.Locals("user_id").(string)

	randStr, _ := gonanoid.New(8)
	res, err := h.CourseRepository.CreateCourse(model.Course{
		Title:        request.Title,
		Description:  request.Description,
		ThumbnailImg: request.ThumbnailImg,
		Slug:         slug.Make(request.Title) + "-" + randStr,
		TeacherID:    teacher,
	})

	if err != nil {
		return c.Status(500).JSON(&http.WebResponse{
			Status:  "fail",
			Message: "Couldn't create course because internal error",
			Data:    nil,
		})
	}

	dataRes := http.CourseHTTP{
		ID:           res.ID,
		Title:        res.Title,
		Description:  res.Description,
		ThumbnailImg: res.ThumbnailImg,
		Slug:         res.Slug,
		PublishedAt:  res.UpdatedAt,
		UpdatedAt:    res.UpdatedAt,
	}

	return c.JSON(&http.WebResponse{
		Status:  "success",
		Message: "Course has been created!",
		Data:    dataRes,
	})
}

func (h *Handlers) GetCourses(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if c.Query("page") == "" || c.Query("limit") == "" {
		return c.Status(400).JSON(&http.WebResponse{
			Status:  "error",
			Message: "Request must specify with query params 'page' and 'limit'",
			Data:    nil,
		})
	}

	pagination, data := h.CourseRepository.FindCourses(page, limit)

	if pagination == nil || data == nil {
		return c.Status(404).JSON(&http.WebResponse{
			Status:  "error",
			Message: "Couldn't find courses data because it's not found!",
			Data:    nil,
		})
	}

	var entries []http.CourseHTTP

	for _, el := range data {
		entries = append(entries, http.CourseHTTP{
			ID:           el.ID,
			Title:        el.Title,
			Slug:         el.Slug,
			Description:  el.Description,
			ThumbnailImg: el.ThumbnailImg,
			PublishedAt:  el.CreatedAt,
			UpdatedAt:    el.UpdatedAt,
		})
	}

	resp := http.CoursePaginationResponse{
		Pagination: pagination,
		Entries:    entries,
	}

	return c.Status(200).JSON(&http.WebResponse{
		Status:  "success",
		Message: "Successfully retrieve Courses!",
		Data:    resp,
	})
}

func (h *Handlers) GetDetailCourse(c *fiber.Ctx) error {
	slug := c.Params("slug")

	if slug == "" {
		return c.Status(400).JSON(&http.WebResponse{
			Status:  "fail",
			Message: "Request must specify 'slug' params!",
			Data:    []interface{}{},
		})
	}

	res, err := h.CourseRepository.FindCourse(map[string]interface{}{
		"slug": slug,
	})

	if err != nil {
		return c.Status(404).JSON(&http.WebResponse{
			Status:  "fail",
			Message: "Couldn't find course because it's not found!",
			Data:    []interface{}{},
		})
	}

	return c.Status(200).JSON(&http.WebResponse{
		Status:  "success",
		Message: "Successfully retrieve detail course!",
		Data: http.CourseHTTP{
			ID:           res.ID,
			Title:        res.Title,
			Slug:         res.Slug,
			ThumbnailImg: res.ThumbnailImg,
			Description:  res.Description,
		},
	})
}

func (h *Handlers) EditCourse(c *fiber.Ctx) error {
	slugParams := c.Params("slug")

	if slugParams == "" {
		return c.Status(400).JSON(&http.WebResponse{
			Status:  "fail",
			Message: "Request must specify 'slug' params",
			Data:    []interface{}{},
		})
	}

	var request http.CourseHTTP
	c.BodyParser(&request)

	randStr, _ := gonanoid.New(8)
	res, err := h.CourseRepository.EditCourse(map[string]interface{}{
		"slug": slugParams,
	}, model.Course{
		Title:        request.Title,
		Description:  request.Description,
		ThumbnailImg: request.ThumbnailImg,
		Slug:         slug.Make(request.Title) + "-" + randStr,
	})

	if err != nil {
		return c.Status(500).JSON(&http.WebResponse{
			Status:  "fail",
			Message: "Couldn't edit course because there's internal error",
			Data:    []interface{}{},
		})
	}

	return c.Status(200).JSON(&http.WebResponse{
		Status:  "success",
		Message: "Course has been edited!",
		Data: http.CourseHTTP{
			ID:           res.ID,
			Title:        res.Title,
			Description:  res.Description,
			Slug:         res.Slug,
			ThumbnailImg: res.ThumbnailImg,
			PublishedAt:  res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
		},
	})
}

func (h *Handlers) DeleteCourse(c *fiber.Ctx) error {
	slugParams := c.Params("slug")

	if slugParams == "" {
		return c.Status(400).JSON(&http.WebResponse{
			Status:  "fail",
			Message: "Request must specify resource using 'slug' params",
			Data:    []interface{}{},
		})
	}

	res, err := h.CourseRepository.DeleteCourse(map[string]interface{}{
		"slug": slugParams,
	})

	if err != nil {
		return c.Status(500).JSON(&http.WebResponse{
			Status:  "fail",
			Message: "Couldn't delete courses because internal error",
			Data:    []interface{}{},
		})
	}

	return c.Status(200).JSON(&http.WebResponse{
		Status:  "success",
		Message: "Course has been deleted!",
		Data: http.CourseHTTP{
			ID:           res.ID,
			Title:        res.Title,
			Description:  res.Description,
			Slug:         res.Slug,
			ThumbnailImg: res.ThumbnailImg,
			PublishedAt:  res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
		},
	})
}

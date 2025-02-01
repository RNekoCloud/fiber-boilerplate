package cmd

import (
	"fmt"
	"net/http"
	"os"

	"api-service/config"
	"api-service/http/handlers"
	"api-service/middleware"
	"api-service/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/urfave/cli"
)

func HTTPGatewayServer(port int) {
	app := fiber.New()

	confDB := config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")

	// Dependency Injection
	newDB := config.NewDBConfig(&confDB)
	courseRepos := repository.NewCourseRepository(newDB)

	// Dependency Injection auth
	authRepos := repository.NewAuthRepository(newDB)

	// Depedency injection User Management
	userManagementRepos := repository.NewUserManagementRepository(newDB)

	handlersDep := handlers.Handlers{
		UserRepository:   authRepos,
		CourseRepository: courseRepos,
		JwtSecret:        []byte(JWT_SECRET),
		Middleware: middleware.Middleware{
			JwtSecret: []byte(JWT_SECRET),
		},
		UserManagement: userManagementRepos,
	}

	// Setup global middleware
	app.Use(logger.New())

	// Setup Course Route
	handlersDep.RouteCourses(app)

	// Setup Auth Route
	handlersDep.RouteAuth(app)

	// Setup User Management Route
	handlersDep.RouterUserManagemet(app)

	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"message": "API Version 1.0.0 is up!",
		})
	})

	app.Listen(fmt.Sprintf(":%d", port))
}

func HTTPGatewayServerCMD() cli.Command {
	return cli.Command{
		Name:  "http-gw-srv",
		Usage: "Run HTTP Gateway Server with specific port",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 8080,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")

			HTTPGatewayServer(port)
			return nil
		},
	}
}

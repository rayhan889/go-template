package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayhan889/go-template/internal/delivery/http"
)

type RouteConfig struct {
	App 			*fiber.App
	UserController 	*http.UserController
	DocsController 	*http.DocsController
	AuthMiddleware 	fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoutes()
	c.SetupAuthRoutes()
}

func (c *RouteConfig) SetupGuestRoutes() {
	c.App.Post("/api/users/register", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)

	c.App.Get("/docs", c.DocsController.Reference)
	c.App.Get("/api/openapi.yaml", c.DocsController.Spec)
}

func (c *RouteConfig) SetupAuthRoutes() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("/api/users/_current", c.UserController.Current)
}
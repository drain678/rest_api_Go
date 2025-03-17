package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/drain678/rest_api_Go.git/app/controller"
)

func ContactRoutes(a *fiber.App) {
	route := a.Group("/api/v1/contact")

	route.Get("/", controller.GetContact)
	route.Post("/", controller.PostContact)
	route.Put("/", controller.PutContact)
	route.Delete("/", controller.DeleteContact)
}
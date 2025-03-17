package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/drain678/rest_api_Go.git/app/controller"
)

func GroupRoutes(a *fiber.App) {
	routes := a.Group("/api/v1/group")

	routes.Get("/", controller.GetGroup)
	routes.Post("/", controller.PostGroup)
	routes.Put("/", controller.PutGroup)
	routes.Delete("/", controller.DeleteGroup)

}
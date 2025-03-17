package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/drain678/rest_api_Go.git/app/model"
)

func PostContact(c *fiber.Ctx) error {
	return c.JSON(model.Contact{})
}

func GetContact(c *fiber.Ctx) error {
	return c.JSON([]model.Contact{})
}

func PutContact(c *fiber.Ctx) error {
	return c.JSON(model.Contact{})

}

func DeleteContact(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
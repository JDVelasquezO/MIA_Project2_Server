package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	println(c)
	return c.JSON(c)
}

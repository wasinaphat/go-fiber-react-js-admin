package controllers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/wasinaphatlilawatthananan/go-admin/database"
	"github.com/wasinaphatlilawatthananan/go-admin/models"
)
func AllPermissions(c *fiber.Ctx) error {
	var  permission[] models.Permission

	database.DB.Find(&permission)

	return c.JSON(permission)
}
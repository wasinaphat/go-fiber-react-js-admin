package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wasinaphatlilawatthananan/go-admin/database"
	"github.com/wasinaphatlilawatthananan/go-admin/models"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

type RoleCreateDTO struct {
	name        string
	permissions []string
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto RoleCreateDTO

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	permissions := make([]models.Permission, len(roleDto.permissions))

	for i, permisionId := range roleDto.permissions {
		id, _ := strconv.Atoi(permisionId)
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name: roleDto.name,
	}

	database.DB.Create(&role)

	return c.JSON(role)

}
func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Find(&role)

	return c.JSON(role)
}
func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}
func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}

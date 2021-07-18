package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, entitiy Enitity, page int) fiber.Map {
	limit := 15
	offset := (page - 1) * limit

	data := entitiy.Take(db, limit, offset)
	total := entitiy.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}

}

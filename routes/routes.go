package routes

import (
	"log"

	"fiber-simple/database"
	"fiber-simple/models"

	"github.com/gofiber/fiber/v2"
)

// CreateCat
// @Summary Create a cat
// @Curl: curl -X POST -H "Content-Type: application/json" -d '{"name":"Garfield","breed":"Persian"}' http://localhost:8080/cats
func CreateCat(c *fiber.Ctx) error {
	db := database.DB.Db
	cat := new(models.Cat)
	if err := c.BodyParser(cat); err != nil {
		log.Fatal(err)
		return err
	}
	db.Create(&cat)
	return c.JSON(cat)
}

// GetCats
// @Summary Get all cats
func GetCats(c *fiber.Ctx) error {
	db := database.DB.Db
	var cats []models.Cat
	db.Find(&cats)
	return c.JSON(cats)
}

// GetCat
// @Summary Get a cat
// @Param id path int true "Cat ID"
func GetCat(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB.Db
	var cat models.Cat
	err := db.First(&cat, id).Error

	if err != nil {
		return c.Status(404).SendString("Cat not found")
	}
	return c.JSON(cat)
}

// SearchCats
// @Summary Search cats
// @Param name query
func SearchCats(c *fiber.Ctx) error {
	query := c.Query("query")
	db := database.DB.Db
	var cats []models.Cat
	db.Where("name LIKE ?", "%"+query+"%").Or("breed LIKE ?", "%"+query+"%").Find(&cats)
	return c.JSON(cats)
}

// DeleteCat
// @Summary Delete a cat
// @Param id path int true "Cat ID"
func DeleteCat(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB.Db
	var cat models.Cat
	err := db.First(&cat, id).Error

	if err != nil {
		return c.Status(404).SendString("Cat not found")
	}

	db.Delete(&cat)
	return c.SendString("Cat successfully deleted")
}

// UpdateCat
// @Summary Update a cat
// @Param id path int true "Cat ID"
func UpdateCat(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB.Db
	var cat models.Cat
	err := db.First(&cat, id).Error

	if err != nil {
		return c.Status(404).SendString("Cat not found")
	}

	if err := c.BodyParser(&cat); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	db.Save(&cat)
	return c.SendString("Cat successfully updated")
}

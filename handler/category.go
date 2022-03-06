package handler

import (
	"errors"
	"github.com/ercancavusoglu/firma-nerede/database"
	"github.com/ercancavusoglu/firma-nerede/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gofiber/fiber/v2"
)

// GetAllCategories query all categories
func GetAllCategories(c *fiber.Ctx) error {
	db := database.DB
	var Categories []model.Category
	db.Find(&Categories)
	return c.JSON(fiber.Map{"status": "success", "message": "All Categories", "data": Categories})
}

// GetCategory query Category
func GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var Category model.Category
	db.Find(&Category, id)

	if Category.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Category found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Category found", "data": Category})
}

// GetCategoryProducts get Category Products on Database
func GetCategoryProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var Category model.Category
	db.Preload(clause.Associations).Find(&Category, id)

	if Category.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Category found with ID", "data": nil})
	}

	//products := getProductByCategory(Category.ID) # It is no longer necessary with Preload
	return c.JSON(fiber.Map{"status": "success", "message": "Products found for this category", "Category": Category.Title, "data": Category.Product})
}

func getProductByCategory(e uint) *model.Product {
	db := database.DB
	var product model.Product
	entity := db.Where(&model.Product{CategoryId: uint(e)}).Find(&product)
	err := entity.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}
	return &product
}

// CreateCategory new Category
func CreateCategory(c *fiber.Ctx) error {
	db := database.DB
	Category := new(model.Category)
	if err := c.BodyParser(Category); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create Category", "data": err})
	}
	db.Create(&Category)
	return c.JSON(fiber.Map{"status": "success", "message": "Created Category", "data": Category})
}

// DeleteCategory delete Category
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var Category model.Category
	db.First(&Category, id)
	if Category.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Category found with ID", "data": nil})

	}
	db.Delete(&Category)
	return c.JSON(fiber.Map{"status": "success", "message": "Category successfully deleted", "data": nil})
}

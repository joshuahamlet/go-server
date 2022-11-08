package cat

import (
	"github.com/go-server/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type Cat struct {
  gorm.Model
  Name string `json:"name"`
  Type string `json:"type"`
  Rating int `json:"rating"`
}

func GetCats(c *fiber.Ctx) error {
  db := database.Connection
  var cats []Cat
  db.Find(&cats)
  return c.JSON(cats)
}

func GetCat(c *fiber.Ctx) error {
  id := c.Params("id")
  db := database.Connection
  var cat Cat
  db.Find(&cat, id)
  return c.JSON(cat)
}


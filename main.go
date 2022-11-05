package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Cat struct {
  gorm.Model
  Name string
  Type string
}

func main() {
app := fiber.New(fiber.Config{ColorScheme : fiber.Colors{Black: "\u001b[92m"}})

  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  
  db.AutoMigrate(&Cat{})

  db.Create(&Cat{Name: "Mr. Pickles", Type: "Tabby"})

  var cat Cat

  db.First(&cat)

  app.Get("/", func(c *fiber.Ctx) error {
    response := fmt.Sprintf("Hi, there!\n\nI have a cat named %s\n\n%s", cat.Name, os.Getenv("SECRET"))
    return c.SendString(response)
  })

  app.Get("/user/+", func(c *fiber.Ctx) error {
      return c.SendString(c.Params("s+"))
  })

  log.Fatal(app.Listen(":3000"))
}

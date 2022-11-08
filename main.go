package main

import (
	"fmt"
	"log"
	"os"

  "github.com/go-server/cat"
	"github.com/go-server/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDataBase() {
  var err error
  database.Connection, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect to database")
  }
  fmt.Println("connected to database")
  database.Connection.AutoMigrate(&cat.Cat{})
  fmt.Println("database migrated")

  database.Connection.Create(&cat.Cat{Name: "Mr. Pickles", Type: "Tabby", Rating: 10})
  fmt.Println("added Mr. Pickles")
}

func initRoutes(app *fiber.App) {
  app.Get("api/v1/cat", cat.GetCats)
  app.Get("api/v1/cat/:id", cat.GetCat)
}

func main() {
  app := fiber.New(fiber.Config{ColorScheme : fiber.Colors{Black: "\u001b[92m"}})
  app.Use(cors.New())
  
  initDataBase()

  initRoutes(app) 
  
  var cat cat.Cat
  database.Connection.First(&cat)

  app.Get("/", func(c *fiber.Ctx) error {
    response := fmt.Sprintf("Hi, there!\n\nI have a cat named %s\n\n%s", cat.Name, os.Getenv("SECRET"))
    return c.SendString(response)
  })

  app.Get("/user/+", func(c *fiber.Ctx) error {
      return c.SendString(c.Params("s+"))
  })

  log.Fatal(app.Listen(":3000"))
}

package main

import (
	"fiber-simple/database"
	"fiber-simple/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	// app.Use(logger.New())
	// app.Get("/monitor", monitor.New())

	app.Get("/", func(c *fiber.Ctx) error {
		asciiCat := `    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)`

		return c.SendString(asciiCat)
	})
	app.Post("/cats", routes.CreateCat)
	app.Get("/cats", routes.GetCats)
	app.Get("/cats/:id", routes.GetCat)
	app.Get("/search", routes.SearchCats)
	app.Delete("/cats/:id", routes.DeleteCat)
	app.Patch("/cats/:id", routes.UpdateCat)

	app.Listen(":8080")
}

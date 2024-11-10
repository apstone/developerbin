package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func main() {
    // Initialize the template engine
    engine := html.New("./templates", ".html")

    // Create a new Fiber app with the template engine
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Serve static files
    app.Static("/static", "./static")

    // Setup a route
    app.Get("/", func(c *fiber.Ctx) error {
        // Pass data to the template (optional)
        data := fiber.Map{
            "Title": "Welcome to Go Fiber with TailwindCSS",
        }
        return c.Render("index", data)
    })

    // Start the server
    app.Listen(":3000")
}


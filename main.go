package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type Post struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Language    string    `gorm:"not null"`
	Timestamp   time.Time `gorm:"autoCreateTime"`
	Content     string    `gorm:"type:text;not null"`
}

var DB *gorm.DB

func initDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB.AutoMigrate(&Post{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		var posts []Post
		DB.Find(&posts)
		return c.Render("index", fiber.Map{"Posts": posts})
	})

	app.Get("/posts/new", func(c *fiber.Ctx) error {
		return c.Render("new", fiber.Map{})
	})

	app.Post("/posts", func(c *fiber.Ctx) error {
		post := new(Post)
		if err := c.BodyParser(post); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		DB.Create(&post)
		return c.Redirect("/")
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var post Post
		if result := DB.First(&post, id); result.Error != nil {
			return c.Status(404).SendString("Post not found")
		}
		return c.Render("show", fiber.Map{"Post": post})
	})

	app.Static("/static", "./static")
}

func main() {
	initDatabase()

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)

	app.Listen(":3000")
}

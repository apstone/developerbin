package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	_ "modernc.org/sqlite"
	"strings"
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
	sqlDB, err := sql.Open("sqlite", "posts.db")
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	DB, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize GORM with database connection: %v", err)
	}

	err = DB.AutoMigrate(&Post{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate the database schema: %v", err)
	}
}

func truncateCode(content string) string {
	maxLines := 5
	lines := strings.Split(content, "\n")
	if len(lines) > maxLines {
		return strings.Join(lines[:maxLines], "\n") + "\n..."
	}
	return content
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
	engine.AddFunc("truncateCode", truncateCode)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)

	app.Listen(":3000")
}

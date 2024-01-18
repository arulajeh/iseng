package main

import (
	"context"
	// "html/template"
	"log"
	"math"
	"time"

	"pertama_go/repositories"
	"pertama_go/types"
	"pertama_go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func connect() {
	clientOptions := options.Client()
	// clientOptions.ApplyURI("mongodb+srv://arulajeh:Bisajadi100@cluster0.cncm26t.mongodb.net")
	clientOptions.ApplyURI("mongodb://arul:Bisajadi100@68.183.232.146")
	var ctx = context.Background()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err.Error())
	}

	db = client.Database("pertama_go")
}

func init() {
	connect()
}

func main() {
	app := fiber.New(
		fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
			Immutable:     true,
			// Prefork:       true,
		},
	)

	// Use CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	// Define the path to your HTML templates
	// templates := template.Must(template.ParseGlob("views/*.html"))
	// app.Static("/", "./public", fiber.Static{
	// 	Compress:      true,
	// 	ByteRange:     true,
	// 	Browse:        true,
	// 	MaxAge:        3600,
	// 	CacheDuration: 10 * time.Minute,
	// })

	// Register the template engine middleware
	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Locals("templates", templates)
	// 	return c.Next()
	// })

	app.Get("/", func(c *fiber.Ctx) error {
		// Render the HTML template
		return c.Render("index.html", fiber.Map{})
	})

	app.Post("/greetings", func(c *fiber.Ctx) error {
		body := new(types.Greetings)
		if err := c.BodyParser(body); err != nil {
			log.Fatal(err.Error())
		}

		body.CreatedAt = time.Now()
		repositories.InsertGreetings(db, body)

		var response = types.Response{
			Status:  201,
			Message: "Success",
		}
		return c.JSON(response)
	})

	app.Get("/greetings", func(c *fiber.Ctx) error {
		pagination := types.Pagination{
			Limit: utils.GetLimit(c),
			Page:  utils.GetPage(c),
		}

		greetings, err := repositories.FindAllGreetings(db, bson.M{}, &pagination)
		if err != nil {
			log.Fatal(err.Error())
		}
		if greetings == nil {
			greetings = []types.Greetings{}
		}

		count := repositories.CountGreetings(db, bson.M{})
		pagination.TotalData = int(count)
		pagination.TotalPage = int(math.Ceil(float64(count) / float64(pagination.Limit)))

		var response = types.Response{
			Status:     200,
			Message:    "Success",
			Data:       greetings,
			Pagination: &pagination,
		}
		return c.JSON(response)
	})
	app.Listen(":3000")
}

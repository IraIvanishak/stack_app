package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/IraIvanishak/stack_app/internal/model"
	"github.com/IraIvanishak/stack_app/internal/repo"
	"github.com/IraIvanishak/stack_app/internal/scrapper"
	"github.com/IraIvanishak/stack_app/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	re := repo.NewJobRepo(client)
	se := services.NewJobService(re)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/scan", func(c *fiber.Ctx) error {
		jobs := scrapper.Scrap()
		return c.JSON(jobs)
	})
	app.Get("/save-test", func(c *fiber.Ctx) error {
		err := re.SaveTest()
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("Test data saved")
	})

	app.Get("/update", func(c *fiber.Ctx) error {
		err := se.FilRaw()
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("Data updated")
	})
	app.Post("/save", func(c *fiber.Ctx) error {
		var data model.Vacancy
		if err := c.BodyParser(&data); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"error":   "cannot parse JSON",
				"details": err.Error(),
			})
		}

		if err := re.Save(data); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error":   "failed to save data",
				"details": err.Error(),
			})
		}

		return c.SendString("Data saved")
	})

	app.Get("/list", func(c *fiber.Ctx) error {
		category := c.Query("category")
		if category == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"error": "category is required",
			})
		}

		data, err := re.GetByCategory(category)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error":   "failed to get data",
				"details": err.Error(),
			})
		}

		return c.JSON(data)
	})

	app.Get("/analytic", func(c *fiber.Ctx) error {
		category := c.Query("category")
		if category == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"error": "category is required",
			})
		}

		data, err := se.GetAnalytic(category)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"error":   "failed to get data",
				"details": err.Error(),
			})
		}

		return c.JSON(data)
	})

	app.Listen(":8080")
}

package main

import (
	"context"
	"fmt"

	"os"

	"github.com/Cutshadows/react-go-crud/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")
	mongo_string := os.Getenv("MONGODB_URI")

	if port == "" {
		port = "2300"
	}

	if mongo_string == "" {
		mongo_string = "mongodb+srv://calendar-user:HfJ6zokjynWQ2Pi8@calendar-db.ynxgqmz.mongodb.net/?retryWrites=true&w=majority"
	}

	app := fiber.New()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_string))

	if err != nil {
		panic(err)
	}
	// coll := client.Database("gomongodb").Collection("users")
	// coll.InsertOne(context.TODO(), bson.D{{
	// 	Key:   "name",
	// 	Value: "douglas",
	// }})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User

		coll := client.Database("gomongodb").Collection("users")
		results, error := coll.Find(context.TODO(), bson.M{})

		if error != nil {
			panic(error)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"users": users,
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user)
		coll := client.Database("gomongodb").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})

	})

	app.Listen(":" + port)
	fmt.Println("Init server on 2300")
}

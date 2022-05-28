package cmd

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"loginUserGo/infrastructure"
	"time"
)

func StartHttp(ctx context.Context, containerDI *infrastructure.ContainerDI) {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	app.Post("/user", containerDI.UserHandler.CreateUser)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

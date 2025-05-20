package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type App struct {
	port     string
	fiberApp *fiber.App
}

func New(fiberApp *fiber.App, port string) *App {
	port = fmt.Sprintf(":%s", port)

	return &App{
		port:     port,
		fiberApp: fiberApp,
	}
}

func (app *App) Run() error {
	return app.fiberApp.Listen(app.port)
}

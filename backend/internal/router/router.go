package router

import (
	"1dv027/wt2/internal/config"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type IoCContainer interface {
	Resolve(name string, lifecycle config.Lifecycle) any
}

type Handler interface {
	Handle(c *fiber.Ctx) error
}

// The main router for the server.
type Router struct {
	container IoCContainer
}

// Creates a new instance of Router.
func NewRouter(container IoCContainer) Router {
	return Router{
		container: container,
	}
}

// Starts the server.
func (r Router) StartRouter() {
	app := fiber.New(fiber.Config{EnableSplittingOnParsers: true})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
	}))

	basePath := os.Getenv("BASE_PATH")
	api := app.Group(basePath + "/api")

	v1 := api.Group("/v1")

	v1.Use(func(c *fiber.Ctx) error {
		authMiddleware := r.container.Resolve("AuthMiddleware", config.Transient).(Handler)
		return authMiddleware.Handle(c)
	})

	v1.Get("/video-games/data", func(c *fiber.Ctx) error {
		dataHandler := r.container.Resolve("GetVideoGamesDataHandler", config.Transient).(Handler)
		return dataHandler.Handle(c)
	})

	v1.Get("/video-games/search", func(c *fiber.Ctx) error {
		searchHandler := r.container.Resolve("GetVideoGamesSearchHandler", config.Transient).(Handler)
		return searchHandler.Handle(c)
	})

	v1.Get("/video-games/parameters", func(c *fiber.Ctx) error {
		parameterHandler := r.container.Resolve("GetVideoGamesParameterHandler", config.Transient).(Handler)
		return parameterHandler.Handle(c)
	})

	absPath, err := filepath.Abs("./internal/public")
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}

	app.Static(basePath+"/", absPath)

	// Serve SPA app for everything that is not explicitly set above.
	app.Get(basePath+"/*", func(c *fiber.Ctx) error {
		log.Print("Trying to get: " + c.OriginalURL())
		return c.SendFile(filepath.Join(absPath, "index.html"))
	})

	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

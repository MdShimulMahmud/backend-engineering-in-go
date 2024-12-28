package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/rewrite"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(limiter.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(pprof.New())
	app.Use(cache.New())
	app.Use(csrf.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(monitor.New())
	app.Use(proxy.New(proxy.Config{
		// Add appropriate configuration here
	}))
	app.Use(requestid.New())
	app.Use(session.New())
	app.Use(rewrite.New())
	app.Use(limiter.New())
	app.Use(pprof.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

package api

import (
	"codeghost/core/agent"
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	app.Post("/analyze", func(c *fiber.Ctx) error {
		diff := c.FormValue("diff")
		log := c.FormValue("log")
		if diff == "" || log == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Diff and log files are required")
		}
		report := agent.Analyze(diff, log)
		if report == "" {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to analyze the files")
		}
		return c.SendString(report)
	})

	app.Listen(":8080")
}
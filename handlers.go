package main

import "github.com/gofiber/fiber/v2"

func index(c *fiber.Ctx) error {
	t_data := fiber.Map{
		"Name": "Matthew",
		"Host": "Meetups Inc.",
	}
	return c.Render("index", t_data)
}

package main

import "github.com/gofiber/fiber/v2"

type Meetup struct {
	Title    string
	Location string
	Date     string
}

func index(c *fiber.Ctx) error {
	Meetups := []Meetup{
		{
			Title:    "Fiesta Floridiana",
			Location: "Fort Lauderdale, FA",
			Date:     "2025-06-13",
		},
		{
			Title:    "Vinyard Reaping",
			Location: "Martha's Vinyard, MA",
			Date:     "2025-04-09",
		},
	}
	return c.Render("index", Meetups)
}

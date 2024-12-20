package main

import (
	"github.com/gofiber/fiber/v2"
)

type Meetup struct {
	Title    string
	Location string
	Date     string
	Image    string
	ImgAlt   string
}

func index(c *fiber.Ctx) error {
	Meetups := []Meetup{
		{
			Title:    "Fiesta Floridiana",
			Location: "Fort Lauderdale, FA",
			Date:     "2025-06-13",
			Image:    "florida-keys.png",
			ImgAlt:   "Florida Keys img",
		},
		{
			Title:    "Vinyard Reaping",
			Location: "Martha's Vinyard, MA",
			Date:     "2025-04-09",
			Image:    "The-Vinyard-SPA.png",
			ImgAlt:   "The Vineyard img",
		},
	}
	return c.Render("index", Meetups)
}

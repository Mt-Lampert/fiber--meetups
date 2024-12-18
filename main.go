package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// defining the Template engine
// templates are files inside `./views` and its subdirectories
// that end with `.go.html`
// the engine will have them all cached!
var engine = html.New("./views", ".go.html")

func main() {
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// access static files
	//
	// ./static/js/htmx.js   => http://localhost:3000/js/htmx.js
	// ./static/css/main.css => http://localhost:3000/css/main.css
	app.Static("/", "./static")

	app.Get("/", index)

	app.Listen(":3000")
}

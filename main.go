// main.go
package main

import (
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type User struct {
	Name    string
	Age     int
	Address string
}

var users = []User{
	{Name: "Alice", Age: 28, Address: "123 Apple St."},
	{Name: "Bob", Age: 35, Address: "456 Berry Blvd."},
	{Name: "Charlie", Age: 22, Address: "789 Cherry Ct."},
}

func main() {
	engine := html.New("/app/views", ".html")

	//app := fiber.New(fiber.Config{
	//	Views: html.New("/app/views", ".html"),
	//})
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/index", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Users": users,
		})
	})

	app.Get("/sort/:column", func(c *fiber.Ctx) error {
		column := c.Params("column")

		switch column {
		case "name":
			sort.Slice(users, func(i, j int) bool {
				return users[i].Name < users[j].Name
			})
		case "age":
			sort.Slice(users, func(i, j int) bool {
				return users[i].Age < users[j].Age
			})
		case "address":
			sort.Slice(users, func(i, j int) bool {
				return users[i].Address < users[j].Address
			})
		}
		return c.Render("index", fiber.Map{
			"Users": users,
		})
	})

	app.Listen(":3200")
}

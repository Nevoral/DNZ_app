package pages

import (
	"BuffetRegister/web/pages/Home"
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func Layout() []lx.Content {
	return []lx.Content{html.DOCTYPE().Html(),
		html.Html(
			html.Head(
				html.Meta().Charset("UTF-8"),
				html.Meta().Name("viewport").Content("width=device-width, initial-scale=1.0"),
				html.Link().Href("/css/output.css").Rel("stylesheet"),
				//html.Link().Href("https://cdn.jsdelivr.net/npm/daisyui@4.10.5/dist/full.min.css").Rel("stylesheet").Type("text/css"),
				//html.Script().Src("https://cdn.tailwindcss.com"),
				html.Script().Src("/js/htmx.min.js"),
				html.Script().Src("/js/register.js"),
				html.Title("Buffet Kasa"),
			),
			html.Body(
				Home.ProductMenuForm(),
			).Class("flex w-full h-full bg-gray-600 justify-center items-center"),
		).Lang("cz"),
	}
}

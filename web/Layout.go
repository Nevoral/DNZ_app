package web

import (
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func Layout(title string, headBlock, bodyBlock, scriptDependence []lx.Content) []lx.Content {
	return []lx.Content{html.DOCTYPE().Html(),
		html.Html(
			html.Head(
				html.Meta().Charset("UTF-8"),
				html.Meta().Name("viewport").Content("width=device-width, initial-scale=1.0"),
				html.Title(title),
				//html.Link().Href("/css/output.css").Rel("stylesheet"),
				//html.Link().Href("https://cdn.jsdelivr.net/npm/daisyui@4.10.5/dist/full.min.css").Rel("stylesheet").Type("text/css"),
				html.Script().Src("https://cdn.tailwindcss.com"),
				headBlock,
			),
			html.Body(
				bodyBlock,
				scriptDependence,
			).Class("flex w-full h-full bg-gray-600 justify-center items-center"),
		).Lang("cz"),
	}
}

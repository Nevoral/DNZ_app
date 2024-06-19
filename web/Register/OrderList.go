package Register

import (
	"fmt"
	hand "github.com/Nevoral/DNZ_app/internal/handlers"
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func OrdList() []lx.Content {
	return []lx.Content{
		html.Div(
			html.H2("Shrnutí Objednávky").Class("font-bold text-lg"),
			html.Ul().Id("summary").Class("my-4 flex flex-col gap-4"),
		),
		html.P("Celkem: 0 Kč").Class("text-lg font-semibold").Id("totalPrice"),
		html.Div(
			html.Button("Vyčistit košík").
				Class("mt-4 px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600").
				Id("clearOrder"),
			html.Button("Zaplaceno").
				Class("mt-4 px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600").
				Id("sendOrder"),
		).Class("flex flex-row justify-between"),
	}
}

func ProductRow(items []hand.Order) (list []lx.Content) {
	for _, item := range items {
		pr := html.Li(
			html.Div(
				html.Span(fmt.Sprintf("%d x", item.Quantity)),
				html.Span(item.Product).Class("mx-2"),
			).Class("flex flex-grow"),
			html.Div(fmt.Sprintf("%d Kč", item.Price*item.Quantity)),
		).Class("flex flex-row justify-between")
		list = append(list, pr)
	}
	return
}

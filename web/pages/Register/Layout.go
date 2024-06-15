package Register

import (
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
	"github.com/Nevoral/LuxeGo/svg"
)

func LayoutLx() lx.Content {
	return html.Div(
		html.Div(
			NavMenu(),
		).Class("flex flex-col w-16 h-full items-center justify-center"),

		html.Div(
			html.Svg(
				svg.Path().D("m0,3v6h11V0H3C1.346,0,0,1.346,0,3Zm9,4H2V3c0-.552.448-1,1-1h6v5Zm4,17h8c1.654,0,3-1.346,3-3v-6h-11v9Zm2-7h7v4c0,.552-.448,1-1,1h-6v-5ZM21,0h-8v13h11V3c0-1.654-1.346-3-3-3Zm1,11h-7V2h6c.552,0,1,.448,1,1v8ZM0,21c0,1.654,1.346,3,3,3h8v-13H0v10Zm2-8h7v9H3c-.552,0-1-.448-1-1v-8Z"),
			).Xmlns("http://www.w3.org/2000/svg").
				ViewBox("0 0 24 24").
				Class("absolute inset-0 m-auto max-w-48 fill-base-100 opacity-50").
				Id("Layer_1").
				Data("name", "Layer_1"),
		).Class("flex-grow text-center relative").
			Id("MenuArea"),
		html.Div(
			OrdList(),
		).Class("w-64 bg-base-100 rounded-l-xl p-4 min-h-64 sticky top-0 flex flex-col justify-between").
			Id("orderSummary"),
	).Class("flex justify-between items-center h-screen")
}

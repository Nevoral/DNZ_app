package Home

import (
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func ProductMenuForm() lx.Content {
	return html.Div(
		html.Div(
			html.H1("Vytvořit nové menu").Class("text-lg font-bold"),
			html.Form(
				html.Label(
					html.Input().Placeholder("Počáteční hodnota Pokladny").Type("number").Name("startRegister").
						Class("input w-full max-w-xs"),
				),
				html.Button("Nové menu").Type("submit").
					Class("btn btn-active btn-primary w-full max-w-xs my-4"),
			).Class("mt-4").
				CustomAtr("hx-post", "/create-menu").
				CustomAtr("hx-target", "#MenuArea").
				CustomAtr("hx-swap", "innerHTML").
				CustomAtr("hx-trigger", "click"),
		).Class("modal-box"),
		html.Form(
			html.Button("close"),
		).Method("dialog").Class("modal-backdrop"),
	).Class("modal h-full justify-center items-center").Id("popupFormMenu")
}

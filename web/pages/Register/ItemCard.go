package Register

import (
	reg "BuffetRegister/internal/database/BuffetRegister"
	"fmt"
	"github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func ItemList(items []reg.Product, id int64, color, category string) LuxeGo.Content {
	cards := make([]LuxeGo.Content, 0, len(items))
	for _, item := range items {
		cards = append(cards, ItemCard(item, color))
	}
	return html.Div(
		cards,
		AddItem(),
		ItemForm(category),
	).Class("basis-7/10 grid grid-cols-6 gap-2 w-full max-w-9xl p-4 overflow-scroll").
		Data("menu-id", fmt.Sprintf("%d", id))
}

func ItemCard(item reg.Product, color string) LuxeGo.Content {
	return html.Div(
		html.Div(
			html.H2(item.Title).Class("card-title text-nowrap truncate"),
			html.Ul(
				html.Li(
					html.Div(
						html.Span("Cena:").Class("font-bold"),
						html.Span(fmt.Sprintf("%d", item.Price)).Class("mx-2").Id("price"),
						html.Span("Kč"),
					).Class("flex flex-grow"),
				).Class("flex flex-row justify-between"),
				html.Li(
					html.Div(
						html.Span("Prodáno:").Class("font-bold"),
						html.Span(fmt.Sprintf("%d", item.Served)).Class("ml-2").
							Id(fmt.Sprintf("count-%d", item.ID)),
						html.Span("x"),
					).Class("flex flex-grow"),
				).Class("flex flex-row justify-between"),
			).Class("my-4 flex flex-col gap-4"),
			html.Div(
				html.Button("Odebrat").Id(fmt.Sprintf("decrementButton-%d", item.ID)).
					Class("btn glass bg-gray-800"),
				//CustomAtr("hx-get", fmt.Sprintf("/decrement?itemId=%d", item.ID)).
				//CustomAtr("hx-target", fmt.Sprintf("#count-%d", item.ID)).
				//CustomAtr("hx-swap", "innerHTML").
				//CustomAtr("hx-trigger", "click"),
			).Class("card-actions justify-end"),
		).Class("card-body justify-between"),
	).Id(fmt.Sprintf("product-%d", item.ID)).Data("id", fmt.Sprintf("%d", item.ID)).
		Data("title", item.Title).Data("price", fmt.Sprintf("%d", item.Price)).
		//CustomAtr("hx-get", fmt.Sprintf("/increment?itemId=%d", item.ID)).
		//CustomAtr("hx-target", fmt.Sprintf("#count-%d", item.ID)).
		//CustomAtr("hx-swap", "innerHTML").
		//CustomAtr("hx-trigger", "click").
		Class(fmt.Sprintf("card h-64 bg-%s text-primary-content shadow-xl hover:scale-105 transition ease-in-out delay-150", color))
}

func AddItem() LuxeGo.Content {
	return html.Div(
		html.Div(
			html.H2("Přidat Produkt").Class("card-title text-nowrap truncate"),
		).Class("card-body justify-center"),
	).Onclick("popupItemForm.showModal()").Class("card h-64 bg-base-100 shadow-xl hover:scale-105 transition ease-in-out delay-150")
}

func ItemForm(category string) LuxeGo.Content {
	var cat LuxeGo.Content
	if category == "food" {
		cat = html.Option("Jidlo").Value("food").Selected()
	} else {
		cat = html.Option("Pití").Value("drink").Selected()
	}
	return html.Dialog(
		html.Div(
			html.H1("Přidat Nový Produkt").Class("text-lg font-bold"),
			html.Form(
				html.Label(
					html.Input().Placeholder("Produkt").Type("Text").Name("title").
						Class("input w-full max-w-xs"),
				),
				html.Label(
					html.Input().Placeholder("Cena").Type("Text").Name("price").
						Class("input w-full max-w-xs"),
				),
				html.Label(
					html.Select(
						cat,
					).Name("category").Class("select w-full max-w-xs"),
				).Class("hidden"),
				html.Button("Add Product").Type("submit").
					Class("btn btn-active btn-primary w-full max-w-xs my-4"),
			).Class("mt-4 py-2").
				CustomAtr("hx-post", "/add-product").
				CustomAtr("hx-target", "#MenuArea").
				CustomAtr("hx-swap", "innerHTML").
				CustomAtr("hx-trigger", "click"),
		).Class("modal-box"),
		html.Form(
			html.Button("close"),
		).Method("dialog").Class("modal-backdrop"),
	).Class("modal").Id("popupItemForm")
}

func Count(msg string) LuxeGo.Content {
	return html.FreeStr(msg)
}

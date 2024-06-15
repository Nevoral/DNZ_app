package Register

import (
	"github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
	"github.com/Nevoral/LuxeGo/svg"
)

func NavMenu() LuxeGo.Content {
	var svgTagClass = "h-10 w-10 p-0.5 stroke-gray-500 fill-gray-300 hover:stroke-teal-500 hover:fill-teal-100 hover:scale-125 hover:-translate-x-0.5"
	var aTagClass = "tooltip tooltip-right flex p-0 my-4 hover:bg-base-100 justify-center"
	return html.Ul(
		html.Li(
			html.A(
				html.Svg(
					svg.Path().D("M9,4H7V0H9Zm4-4H11V4h2ZM5,0H3V4H5ZM24,18c0,5.221-9.4,6-15,6A81.716,81.716,0,0,1,.835,22.986L0,22.847V16A9.01,9.01,0,0,1,9,7a18.144,18.144,0,0,1,2.409.164A17.517,17.517,0,0,1,17.091,6l.434-.014.538.537,1.716-1.715a1.5,1.5,0,1,1,2.063-.65A1.485,1.485,0,0,1,22.5,4a1.5,1.5,0,1,1-1.307,2.222L19.477,7.937l.538.538L20,8.909c-.006.19-.04.9-.164,1.833C22.442,12.856,24,15.548,24,18ZM10,12.312a3.667,3.667,0,0,0,1.08,2.608,3.777,3.777,0,0,0,5.215,0c1.145-1.146,1.577-4.153,1.683-5.653L16.733,8.022c-1.5.106-4.507.537-5.653,1.683A3.667,3.667,0,0,0,10,12.312ZM22,18a7.516,7.516,0,0,0-2.613-4.988,7.543,7.543,0,0,1-1.678,3.322A5.678,5.678,0,0,1,9.081,9,7.075,7.075,0,0,0,2,16v5.149A68.5,68.5,0,0,0,9,22C15,22,22,20.953,22,18Z"),
				).Xmlns("http://www.w3.org/2000/svg").
					ViewBox("0 0 24 24").
					Width("512").
					Height("512").
					Data("name", "layer 1").
					Id("Layer_1").
					Class(svgTagClass),
			).Class(aTagClass).
				Data("type", "Jídlo").
				CustomAtr("hx-get", "/food").
				CustomAtr("hx-target", "#MenuArea").
				CustomAtr("hx-swap", "innerHTML").
				CustomAtr("hx-trigger", "click"),
		),
		html.Li(
			html.A(
				html.Svg(
					svg.Path().D("M11.026,4h-2V0h2Zm4-4h-2V4h2ZM0,24H20V22H0ZM23.993,13.143A3.981,3.981,0,0,1,20,17H16.845a22.085,22.085,0,0,1-2.431,3H5.59A21.943,21.943,0,0,1,.033,9.4,2.844,2.844,0,0,1,.674,7.087,3.047,3.047,0,0,1,3.007,6h0L17,6A3.05,3.05,0,0,1,19.328,7.09,2.84,2.84,0,0,1,19.97,9.4c-.035.2-.081.4-.123.6H20C22.462,10,23.993,11.205,23.993,13.143ZM17.8,8.38A1.061,1.061,0,0,0,17,8L3.008,8h0a1.065,1.065,0,0,0-.8.376.841.841,0,0,0-.2.685A19.193,19.193,0,0,0,6.44,18h7.124A19.181,19.181,0,0,0,18,9.064.841.841,0,0,0,17.8,8.38Zm4.2,4.763c0-.758-.672-1.143-2-1.143h-.687a20.161,20.161,0,0,1-1.279,3H20A1.982,1.982,0,0,0,21.994,13.143ZM7.026,0h-2V4h2Z"),
				).Xmlns("http://www.w3.org/2000/svg").
					ViewBox("0 0 24 24").
					Width("512").
					Height("512").
					Data("name", "layer 1").
					Id("Layer_1").
					Class(svgTagClass),
			).Class(aTagClass).
				Data("type", "Pití").
				CustomAtr("hx-get", "/drink").
				CustomAtr("hx-target", "#MenuArea").
				CustomAtr("hx-swap", "innerHTML").
				CustomAtr("hx-trigger", "click"),
		),
		html.Li(
			html.A(
				html.Svg(
					svg.Path().D("m12,12h4.242l6.879-6.879c1.17-1.17,1.17-3.072,0-4.242s-3.072-1.17-4.242,0l-6.879,6.879v4.242Zm2-3.414l6.293-6.293c.391-.391,1.023-.391,1.414,0s.39,1.024,0,1.414l-6.293,6.293h-1.414v-1.414Zm6,8.414v-5.93l-2,2v3.93h-8v3.5c0,.827-.673,1.5-1.5,1.5s-1.5-.673-1.5-1.5V3.5c0-.539-.133-1.044-.351-1.5h8.281L16.89.039c-.13-.015-.257-.039-.39-.039H3.5C1.57,0,0,1.57,0,3.5v3.5h5v13.5c0,1.93,1.57,3.5,3.5,3.5h12c1.93,0,3.5-1.57,3.5-3.5v-3.5h-4ZM5,5h-3v-1.5c0-.827.673-1.5,1.5-1.5s1.5.673,1.5,1.5v1.5Zm17,15.5c0,.827-.673,1.5-1.5,1.5h-8.838c.217-.455.338-.964.338-1.5v-1.5h10v1.5Z"),
				).Xmlns("http://www.w3.org/2000/svg").
					ViewBox("0 0 24 24").
					Width("512").
					Height("512").
					Data("name", "layer 1").
					Id("Layer_1").
					Class(svgTagClass),
			).Class(aTagClass).
				Data("type", "Otevřené Účty").
				CustomAtr("hx-get", "/bills").
				CustomAtr("hx-target", "#MenuArea").
				CustomAtr("hx-swap", "innerHTML").
				CustomAtr("hx-trigger", "click"),
		).Class("hidden"),
		html.Li(
			html.A(
				html.Svg(
					svg.Path().D("m7.456,22h16.515l-3.075-17.424c-.263-1.493-1.554-2.576-3.07-2.576H6.148c-1.516,0-2.807,1.083-3.07,2.576L.003,22h5.198l-.333-2h-2.481l2.358-13.361,2.711,15.361ZM17.826,4c.543,0,1.006.388,1.101.923l2.661,15.077h-12.453L6.31,4h11.516Zm-8.249,7h8.391l.353,2h-8.391l-.353-2Zm-.353-2l-.353-2h8.391l.353,2h-8.391Zm1.059,6h8.391l.353,2h-8.391l-.353-2Z"),
				).Xmlns("http://www.w3.org/2000/svg").
					ViewBox("0 0 24 24").
					Width("512").
					Height("512").
					Data("name", "layer 1").
					Id("Layer_1").
					Class(svgTagClass),
			).Class(aTagClass).
				Onclick("popupFormMenu.showModal()").
				Data("type", "Nové Menu").
				CustomAtr("hx-get", "/menu").
				CustomAtr("hx-target", "#popupContent").
				CustomAtr("hx-swap", "innerHTML").
				CustomAtr("hx-trigger", "click"),
		),
		html.Li(
			html.Label(
				html.Input().Type("checkbox").Value("synthwave").Class("theme-controller hidden"),
				html.Svg(
					svg.Path().D("M5.64,17l-.71.71a1,1,0,0,0,0,1.41,1,1,0,0,0,1.41,0l.71-.71A1,1,0,0,0,5.64,17ZM5,12a1,1,0,0,0-1-1H3a1,1,0,0,0,0,2H4A1,1,0,0,0,5,12Zm7-7a1,1,0,0,0,1-1V3a1,1,0,0,0-2,0V4A1,1,0,0,0,12,5ZM5.64,7.05a1,1,0,0,0,.7.29,1,1,0,0,0,.71-.29,1,1,0,0,0,0-1.41l-.71-.71A1,1,0,0,0,4.93,6.34Zm12,.29a1,1,0,0,0,.7-.29l.71-.71a1,1,0,1,0-1.41-1.41L17,5.64a1,1,0,0,0,0,1.41A1,1,0,0,0,17.66,7.34ZM21,11H20a1,1,0,0,0,0,2h1a1,1,0,0,0,0-2Zm-9,8a1,1,0,0,0-1,1v1a1,1,0,0,0,2,0V20A1,1,0,0,0,12,19ZM18.36,17A1,1,0,0,0,17,18.36l.71.71a1,1,0,0,0,1.41,0,1,1,0,0,0,0-1.41ZM12,6.5A5.5,5.5,0,1,0,17.5,12,5.51,5.51,0,0,0,12,6.5Zm0,9A3.5,3.5,0,1,1,15.5,12,3.5,3.5,0,0,1,12,15.5Z"),
				).Xmlns("http://www.w3.org/2000/svg").
					ViewBox("0 0 24 24").
					Class("swap-off fill-current w-10 h-10"),
				html.Svg(
					svg.Path().D("M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Zm-9.5,6.69A8.14,8.14,0,0,1,7.08,5.22v.27A10.15,10.15,0,0,0,17.22,15.63a9.79,9.79,0,0,0,2.1-.22A8.11,8.11,0,0,1,12.14,19.73Z"),
				).Xmlns("http://www.w3.org/2000/svg").
					ViewBox("0 0 24 24").
					Class("swap-off fill-current w-10 h-10"),
			).Class("swap swap-rotate"),
		).Class("hidden"),
	).Class("menu bg-base-100 w-full rounded-r-xl p-0 m-0")
}

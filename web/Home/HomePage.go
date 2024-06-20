package Home

import (
	"github.com/Nevoral/DNZ_app/web"
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func Page() []lx.Content {
	return web.Layout(
		"Dvůr nové Zámky",
		[]lx.Content{html.Link().Href("http://localhost:8070/css?path=/Home&name=transitions.css").
			Rel("stylesheet").
			Type("text/tailwindcss")},
		[]lx.Content{
			html.Button("Přihlásit se").Popovertarget("popupform").Popovertargetaction("show").
				Class("px-4 py-2 bg-blue-500 text-white rounded"),
			html.Div().Id("ToastContainer"),
			html.Div(
				PopupWindowCon(
					WelcomeTab(true),
					AuthTab(true),
				),
			).Class("static size-2/3 p-2 rounded-2xl shadow-2xl bg-gray-200 bg-opacity-50 backdrop-blur-sm").
				Id("popupform").Popover(),
		},
		[]lx.Content{
			html.Script().Src("http://localhost:8070/js?name=htmx.min.js"),
			html.Script().Src("http://localhost:8070/js?path=/Home&name=validatePassword.js"),
			html.Script().Src("http://localhost:8070/js?path=/Home&name=validateUsername.js"),
			html.Script().Src("http://localhost:8070/js?path=/Home&name=validateEmail.js"),
			html.Script().Src("http://localhost:8070/js?path=/Home&name=validation.js"),
		},
	)
}

func PopupWindowCon(content ...lx.Content) lx.Content {
	return html.Div(
		content,
	).Class("flex flex-row size-full items-start justify-start bg-gray-100 rounded-xl")
}

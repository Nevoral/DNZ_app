package Home

import (
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func LandingPage() []lx.Content {
	return []lx.Content{html.DOCTYPE().Html(),
		html.Html(
			html.Head(
				html.Meta().Charset("UTF-8"),
				html.Meta().Name("viewport").Content("width=device-width, initial-scale=1.0"),
				html.Link().Href("/css/output.css").Rel("stylesheet"),
				//html.Link().Href("https://cdn.jsdelivr.net/npm/daisyui@4.10.5/dist/full.min.css").Rel("stylesheet").Type("text/css"),
				html.Script().Src("https://cdn.tailwindcss.com"),
				html.Script().Src("http://localhost:8080/js?name=htmx.min.js"),
				html.Title("Sign Up"),
				html.Style(`::backdrop {backdrop-filter: blur(3px);}`),
			),
			html.Body(
				html.Button("Zaregistrovat").Popovertarget("signupform").Popovertargetaction("show").
					Class("px-4 py-2 bg-blue-500 text-white rounded"),
				html.Div().Id("ToastContainer"),
				html.Div(
					html.Div(
						SignUpTab(),
						html.Div().Class("box-border w-96 items-start justify-start"),
					).Class("flex flex-row max-w-4xl items-start justify-start bg-gray-100 rounded-xl bg-cover bg-no-repeat bg-center bg-[url(http://localhost:8080/assets?path=/Home&name=spolek.png)]"),
				).Class("static p-2 rounded-2xl shadow-2xl bg-gray-200 bg-opacity-50 backdrop-blur-sm").Id("signupform").Popover(),
				html.Script().Src("http://localhost:8080/js?path=/Home&name=validatePassword.js"),
				html.Script().Src("http://localhost:8080/js?path=/Home&name=validateUsername.js"),
				html.Script().Src("http://localhost:8080/js?path=/Home&name=validateEmail.js"),
				html.Script().Src("http://localhost:8080/js?path=/Home&name=validation.js"),
			).Class("bg-gray-200 font-mono min-h-screen"),
		).Lang("en"),
	}
}

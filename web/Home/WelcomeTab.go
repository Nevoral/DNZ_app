package Home

import (
	"fmt"
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func WelcomeTab(isLogIn bool) lx.Content {
	position := "right-2"
	containerRounded := "rounded-r-xl"
	msg := "už máš"
	redirectButton := RedirectBtn("Přihlásit", "/logintab")
	bgImage := "path=/Home&name=dnz_haf"

	if isLogIn {
		position = "left-2"
		msg = "nemáš"
		containerRounded = "rounded-l-xl"
		redirectButton = RedirectBtn("Registrovat", "/signuptab")
		bgImage = "path=/Home&name=spolek.jpg"
	}
	return html.Div(
		html.Div(
			html.Img().Src("http://localhost:8070/assets?path=/Home&name=dnz_logo.svg").
				Class(fmt.Sprintf("absolute top-2 %s h-16 w-16 rounded-xl", position)),
			html.Div(
				html.P("Vítejte na Nových Zámcích").
					Class("text-4xl text-black text-center text-pretty font-bold mb-4"),
				html.P(fmt.Sprintf("Pokud %s založený účet klikni na tlačítko dole a pokračuj do aplikace.", msg)).
					Class("mx-8 mb-4 text-lg text-neutral-700 text-center text-pretty font-bold mb-4"),
				redirectButton,
			).Class("flex flex-col size-full justify-center items-center"),
		).Class("relative size-full bg-white bg-opacity-50 rounded-xl backdrop-blur-sm p-4"),
	).Class(fmt.Sprintf("box-border size-full p-5 %s bg-cover bg-no-repeat bg-center bg-[url(http://localhost:8070/assets?%s)]", containerRounded, bgImage))
}

func RedirectBtn(msg, urlTarget string) lx.Content {
	return html.Button(msg).Type("button").
		Class("w-72 px-4 py-2 z-0 bg-gradient-to-r from-cyan-600 to-green-600 text-white font-bold hover:scale-105 transition ease-in-out delay-150 shadow-xl rounded-full mt-1 mb-4").
		CustomAtr("hx-get", urlTarget).
		CustomAtr("hx-target", "#popupform").
		CustomAtr("hx-swap", "innerHTML transition:true")
}

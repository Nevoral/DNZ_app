package Home

import (
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
)

func EmailConfirm(email string) lx.Content {
	return html.Div(
		html.H2("Už jen poslední krok...").Class("text-4xl text-center text-pretty font-bold mb-4"),
		html.P("Prosím ověřte svoji E-mailovou adresu, abyste se mohli přihlásit. Kdybyste někdy zapomněli heslo můžeme vám poslat reset starého hesla. Poslali jsme potvrzovací E-mail na adresu.").
			Class("text-center text-neutral-800 mb-4"),
		html.P(email).Class("font-bold text-black text-lg m-2"),
		html.Div(
			html.P("Pokud vám nic nepřišlo "),
			html.A("klikněte sem.").Class("font-bold text-cyan-600 px-2"),
		).Class("flex flex-row"),
	).Class("flex flex-col justify-center items-center size-full p-8 bg-white text-black shadow-lg rounded-lg")
}

package Home

import (
	"fmt"
	lx "github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
	"github.com/Nevoral/LuxeGo/svg"
)

func AuthTab(isLogin bool) lx.Content {
	return html.Div(
		html.H2("Tam kde podkovy zvoní...").Class("text-4xl text-center text-pretty font-bold mb-4"),
		html.Div(
			CircleButton(
				html.Svg(
					"<g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><path d=\"M273.81 52.973C313.806.257 369.41 0 369.41 0s8.271 49.562-31.463 97.306c-42.426 50.98-90.649 42.638-90.649 42.638s-9.055-40.094 26.512-86.971zM252.385 174.662c20.576 0 58.764-28.284 108.471-28.284 85.562 0 119.222 60.883 119.222 60.883s-65.833 33.659-65.833 115.331c0 92.133 82.01 123.885 82.01 123.885s-57.328 161.357-134.762 161.357c-35.565 0-63.215-23.967-100.688-23.967-38.188 0-76.084 24.861-100.766 24.861C89.33 608.73 0 455.666 0 332.628c0-121.052 75.612-184.554 146.533-184.554 46.105 0 81.883 26.588 105.852 26.588z\" fill=\"#999\"></path></g>",
				).Width("20px").Height("20px").ViewBox("-56.24 0 608.728 608.728").Xmlns("http://www.w3.org/2000/svg").PreserveAspectRatio("xMidYMid").
					Fill("#000000"),
			),
			CircleButton(
				html.Svg(
					"<g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><path d=\"M255.878 133.451c0-10.734-.871-18.567-2.756-26.69H130.55v48.448h71.947c-1.45 12.04-9.283 30.172-26.69 42.356l-.244 1.622 38.755 30.023 2.685.268c24.659-22.774 38.875-56.282 38.875-96.027\" fill=\"#4285F4\"></path><path d=\"M130.55 261.1c35.248 0 64.839-11.605 86.453-31.622l-41.196-31.913c-11.024 7.688-25.82 13.055-45.257 13.055-34.523 0-63.824-22.773-74.269-54.25l-1.531.13-40.298 31.187-.527 1.465C35.393 231.798 79.49 261.1 130.55 261.1\" fill=\"#34A853\"></path><path d=\"M56.281 156.37c-2.756-8.123-4.351-16.827-4.351-25.82 0-8.994 1.595-17.697 4.206-25.82l-.073-1.73L15.26 71.312l-1.335.635C5.077 89.644 0 109.517 0 130.55s5.077 40.905 13.925 58.602l42.356-32.782\" fill=\"#FBBC05\"></path><path d=\"M130.55 50.479c24.514 0 41.05 10.589 50.479 19.438l36.844-35.974C195.245 12.91 165.798 0 130.55 0 79.49 0 35.393 29.301 13.925 71.947l42.211 32.783c10.59-31.477 39.891-54.251 74.414-54.251\" fill=\"#EB4335\"></path></g>",
				).Width("20px").Height("20px").ViewBox("-3 0 262 262").Xmlns("http://www.w3.org/2000/svg").PreserveAspectRatio("xMidYMid").
					Fill("#000000"),
			),
			CircleButton(
				html.Svg(
					"<g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><circle cx=\"420.945\" cy=\"296.781\" r=\"294.5\" fill=\"#3c5a9a\"></circle><path d=\"M516.704 92.677h-65.239c-38.715 0-81.777 16.283-81.777 72.402.189 19.554 0 38.281 0 59.357H324.9v71.271h46.174v205.177h84.847V294.353h56.002l5.067-70.117h-62.531s.14-31.191 0-40.249c0-22.177 23.076-20.907 24.464-20.907 10.981 0 32.332.032 37.813 0V92.677h-.032z\" fill=\"#ffffff\"></path></g>",
				).Width("20px").Height("20px").ViewBox("126.445 2.281 589 589").Xmlns("http://www.w3.org/2000/svg").PreserveAspectRatio("xMidYMid").
					Fill("#000000"),
			),
		).Class("flex flex-row space-x-4 justify-center mb-3"),
		html.Div("nebo").Class("text-center text-neutral-800 mb-4"),
		AuthForm(isLogin),
	).Id("authTab").Class("flex flex-col justify-center items-center size-full p-8 bg-white text-black shadow-lg rounded-lg")
}

func AuthForm(isLogin bool) lx.Content {
	var (
		inputs = []lx.Content{
			InputComponent("text", "username", "Uživatelské jméno", html.Svg().Class("hidden")),
			InputComponent("email", "email", "E-mailová adresa", html.Svg().Class("hidden")),
			InputComponent("password", "password", "Heslo",
				html.Svg(
					svg.Defs().Id("defs1"),
					svg.G(
						svg.Path().FillRule("evenodd").ClipRule("evenodd").D("M 6.30147,15.5771 C 4.77832,14.2684 3.6904,12.7726 3.18002,12 3.6904,11.2274 4.77832,9.73158 6.30147,8.42294 7.87402,7.07185 9.81574,6 12,6 14.1843,6 16.1261,7.07185 17.6986,8.42294 19.2218,9.73158 20.3097,11.2274 20.8201,12 20.3097,12.7726 19.2218,14.2684 17.6986,15.5771 16.1261,16.9282 14.1843,18 12,18 9.81574,18 7.87402,16.9282 6.30147,15.5771 Z M 12,4 C 9.14754,4 6.75717,5.39462 4.99812,6.90595 3.23268,8.42276 2.00757,10.1376 1.46387,10.9698 c -0.41081,0.6287 -0.41081,1.4317 0,2.0604 0.5437,0.8322 1.76881,2.547 3.53425,4.0639 C 6.75717,18.6054 9.14754,20 12,20 c 2.8525,0 5.2429,-1.3946 7.002,-2.9059 1.7654,-1.5169 2.9905,-3.2317 3.5342,-4.0639 0.4108,-0.6287 0.4108,-1.4317 0,-2.0604 C 21.9925,10.1376 20.7674,8.42276 19.002,6.90595 17.2429,5.39462 14.8525,4 12,4 Z m -2,8 c 0,-1.1046 0.8955,-2 2,-2 1.1046,0 2,0.8954 2,2 0,1.1046 -0.8954,2 -2,2 -1.1045,0 -2,-0.8954 -2,-2 z m 2,-4 c -2.2091,0 -3.99996,1.79086 -3.99996,4 0,2.2091 1.79086,4 3.99996,4 2.2092,0 4,-1.7909 4,-4 0,-2.20914 -1.7908,-4 -4,-4 z").
							CustomAtr("fill", "#000000").Id("path1"),
					).Id("SVGRepo_iconCarrier").CustomAtr("transform", "translate(-1.1557626,-1.4339278)"),
					svg.Path().D("m 1.557854,2.1303873 c 0.00488,1.4830474 15.052021,18.3226217 16.319917,18.3047237 1.563225,0.0073 2.271734,-0.558679 2.200115,-2.182343 C 19.992579,17.20951 5.6340458,0.03223893 3.582136,0.05672123 2.1420594,0.07479593 1.573825,0.63381353 1.557854,2.1303873 Z").Style("fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:0.113386").
						Id("eyeSlash"),
				).Width("20px").Height("20px").ViewBox("0 0 24 24").Xmlns("http://www.w3.org/2000/svg").
					Fill("none").Onclick(`togglePasswordVisibility();`).Class("pl-1.5"),
			),
		}
		checkbox = []lx.Content{
			html.Input().Type("checkbox").Name("terms").Class("bg-green-500").Id("terms"),
			html.Span("Souhlasím s").Class("px-2 text-gray-800 text-xs"),
			html.A("Podmínkami a služby").Class("text-gray-800 text-xs hover:font-bold"),
		}
		button = html.Button("Registrovat").Disabled().Type("button").Id("registerBtn").
			Class("w-72 px-4 py-2 z-0 bg-gradient-to-r from-cyan-600 to-green-600 text-white font-bold transition ease-in-out delay-150 shadow-xl rounded-full mt-1 mb-4").
			CustomAtr("hx-post", "/signup").
			CustomAtr("hx-target", "#popupform").
			CustomAtr("hx-swap", "innerHTML")
		formId = "signupForm"
	)

	if isLogin {
		inputs = inputs[1:]
		checkbox = logInCheckbox()
		button = logInButton()
		formId = "loginForm"
	}

	return html.Form(
		inputs,
		html.Div(
			checkbox,
		).Class("flex flex-row items-start px-5 pb-2 justify-start "),
		button,
	).Id(formId).Class("flex flex-col mx-8")
}

func logInCheckbox() []lx.Content {
	return []lx.Content{
		html.Input().Type("checkbox").Class("bg-green-500"),
		html.Span("Zůstat přihlášen").Class("px-2 text-gray-800 text-xs"),
	}
}

func logInButton() *html.ComponentHtmlTag {
	return html.Button("Přihlásit").Disabled().Type("button").Id("loginBtn").
		Class("w-72 px-4 py-2 z-0 bg-gradient-to-r from-cyan-600 to-green-600 text-white font-bold transition ease-in-out delay-150 shadow-xl rounded-full mt-1 mb-4").
		CustomAtr("hx-post", "/login").
		CustomAtr("hx-target", "#authTab").
		CustomAtr("hx-swap", "innerHTML")
}

func InputComponent(inputType, id, placeholder string, element lx.Content) lx.Content {
	return html.Div(
		html.Div(placeholder).Class("text-gray-800 text-xs font-bold mx-4"),
		html.Div(
			html.Div(
				html.Input().Type(inputType).Name(id).Placeholder(placeholder).Id(id).
					Onfocus(fmt.Sprintf("addValidationListeners(this, %sValidate)", id)).
					Class("w-full text-base focus:outline-none bg-white bg-opacity-0"),
				element,
			).Class("flex flex-row justify-between items-center mx-4 w-full"),
		).Id(id+"Shadow").
			Class("flex flex-row shadow-lg text-neutral-800 rounded-full pb-1 w-full h-10 bg-green-50 hover:bg-green-100 transition ease-in-out delay-150s"),
		html.Span("").Id(id+"MsgBox").
			Class("relative w-auto text-neutral-800 p-1 px-2 rounded-xl bg-red-500 bg-opacity-0 text-xs mx-4 font-bold -top-2"),
	).Class("flex flex-col justify-center items-start mb-2 w-72")
}

func CircleButton(image lx.Content) lx.Content {
	return html.Button(image).
		Class("bg-green-50 hover:bg-green-100 hover:scale-125 transition ease-in-out delay-150 rounded-full p-4 shadow-lg")
}

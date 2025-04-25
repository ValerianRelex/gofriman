package main

import (
	"io"
	"net/http"
	"strings"
)

type StringHandler struct {
	message string
}

var counter int

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Method: %v", request.Method)
	Printfln("URL: %v", request.URL)
	Printfln("HTTP Version: %v", request.Proto)
	Printfln("Host: %v", request.Host)
	for name, val := range request.Header {
		Printfln("Header: %v, Value: %v", name, val)
	}
	counter++
	Printfln("--- %v ---", counter)

	// if request.URL.Path == "/favicon.ico" {
	// 	Printfln("Request for icon detected - returning 404")
	// 	writer.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	Printfln("Request for %v", request.URL.Path)

	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/message":
		io.WriteString(writer, sh.message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
}

func HTTPSRedirect(writer http.ResponseWriter, request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5500" + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
	// err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
	
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	
	// статический HTTP сервер
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))
	
	// если здесь обработчик = nil, тогда юзаем http.Handle()
	err := http.ListenAndServe(":5000", nil)

	// не повторял до конца, но понял, что http.HandlerFunc() это АДАПТЕР, представляет интерфейс http.Handler
 	http.Handle("/templates/", http.StripPrefix("/templates/", http.HandlerFunc(HTTPSRedirect)))

	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer", "certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()
	
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

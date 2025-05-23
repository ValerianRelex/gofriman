package main

import (
	_ "fmt"
	"io"
	"net/http"
)

// func mainPage(res http.ResponseWriter, req *http.Request) {
// 	body := fmt.Sprintf("Method: %s\r\n", req.Method)
// 	body += "Header ===============\r\n"
// 	for k, v := range req.Header {
// 		body += fmt.Sprintf("%s: %v\r\n", k, v)
// 	}
// 	body += "Query parameters ===============\r\n"
// 	for k, v := range req.URL.Query() {
// 		body += fmt.Sprintf("%s: %v\r\n", k, v)
// 	}
// 	res.Write([]byte(body))
// }

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Auth(login, password string) bool {
	return login == `guest` && password == `demo`
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Добро пожаловать!")
		} else {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		}
		return
	} else {
		io.WriteString(w, form)
	}
}

func main() {

	// mux := http.NewServeMux()
	// mux.HandleFunc(`/`, mainPage)

	// err := http.ListenAndServe(`:8080`, mux)
	// if err != nil {
	// 	panic(err)
	// }

	err := http.ListenAndServe(":8080", http.HandlerFunc(mainPage))
	if err != nil {
		panic(err)
	}
}

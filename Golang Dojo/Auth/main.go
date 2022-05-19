package main

import (
	"html/template"
	"net/http"
)

func getSignInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-in.html", nil)
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-up.html", nil)
}

func templating(W http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(W, fileName, data)
}

func signInUser(w http.ResponseWriter, r *http.Request) {

}

func signUpUser(w http.ResponseWriter, r *http.Request) {

}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}

}

func main() {
	http.HandleFunc("/", userHandler)
	http.Handle("/css", http.FileServer(http.Dir("css/")))
	http.ListenAndServe("", nil)
}

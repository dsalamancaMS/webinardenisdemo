package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageVars struct {
	Message  string
	Language string
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func render(w http.ResponseWriter, tmpl string, pageVars PageVars) {

	tmpl = fmt.Sprintf("views/%s", tmpl)
	t, err := template.ParseFiles(tmpl)

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, pageVars) //execute the template and pass in the variables to fill the gaps

	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Message:  "Hola Rocket Girls!",
		Language: "Go Lang",
	}
	render(w, "index.html", pageVars)
}

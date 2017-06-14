package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type dataSender struct {
	username string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()               // parse arguments, you have to call this by yourself
	fmt.Println("Form", r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	data := &dataSender{}
	data.username = "World!"

	// tmpl is the HTML template that drives the user interface.
	var tmpl = template.Must(template.New("tmpl").Parse(`
	<!DOCTYPE html><html><body><center>
		<h1>Hello {.username}</h1>
	</center></body></html>
	`))

	err := tmpl.Execute(w, r)
	if err != nil {
		log.Print(err)
	}

	//fmt.Fprintf(w, "Hello World!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

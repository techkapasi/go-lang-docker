package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

func runHTMLTemplate(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()  // parse arguments, you have to call this by yourself

	data := struct {
		Name     string
		Company  string
		URLLong [] string
	}{
		"Hasan Kapasi",
		"10Pearls",
		r.Form["url_long"],
	}

	var tmpl = template.Must(template.New("tmpl").Parse(`
	<!DOCTYPE html>
	<html>
		<body>
			<center>
				<h1>Hello From {{ .Name }} ~ {{ .Company }}</h1>
				<br><br>
				URLLong : {{ .URLLong }}
			</center>
		</body>
	</html>
	`))

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		username, password := r.Form["username"], r.Form["password"]

		fmt.Fprintf(w, "username: %s \n", username[0])
		fmt.Fprintf(w, "password: %s \n", password[0])
	}
}

func main() {

	//Routing Rules
	http.HandleFunc("/", runHTMLTemplate)
	http.HandleFunc("/login", login)

	// Setting listen port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

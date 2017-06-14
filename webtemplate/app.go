package main

import (
	"html/template"
	"log"
	"net/http"
)

func runHTMLTemplate(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Name     string
		Company  string
	}{
		"Hasan Kapasi",
		"10Pearls",
	}

	var tmpl = template.Must(template.New("tmpl").Parse(`
	<!DOCTYPE html>
	<html>
		<body>
			<center>
				<h1>Hello From {{ .Name }} ~ {{ .Company }}</h1>
			</center>
		</body>
	</html>
	`))


	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", runHTMLTemplate2)       // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

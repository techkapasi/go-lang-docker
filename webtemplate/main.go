package main

import (
	"html/template"
	"log"
	"net/http"
)


func runHTMLTemplate2(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name     string
		Company  string
	}{
		"Asjad",
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


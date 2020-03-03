package gui

import (
	"html/template"
	"net/http"
)

func choiceEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	model := make(map[string]string)

	tmpl, err := template.New("choice").Parse(choicePageTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(writer, model)
}

func cssEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/css")

	writer.Write([]byte(css))
}

func GuiServer() {
	http.HandleFunc("/", choiceEndpoint)
	http.HandleFunc("/style.css", cssEndpoint)

	http.ListenAndServe(":8080", nil)
}

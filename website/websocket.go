package website

import (
	"html/template"
	"net/http"
)

func websocketHandle() func(http.ResponseWriter, *http.Request) {
	tmpl, err := template.ParseFiles("website/templates/websocket.html")
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {

		tmpl.Execute(w, nil)
	}

}

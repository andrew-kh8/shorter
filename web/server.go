package web

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

type dashboardParams struct {
	URLCount int
}

type listParams struct {
	URLs []URLShort
}

type URLShort struct {
	Name     string
	URL      string
	URLShort string
}

//go:embed *.html
var templates embed.FS

func dashboard(w http.ResponseWriter, r *http.Request) {
	log.Print("Dash")
	params := dashboardParams{
		URLCount: 3,
	}

	parsed, err := template.ParseFS(templates, "dashboard.html")
	if err != nil {
		log.Print("ParseFS Failed", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	parsed.Execute(w, params)
}

func list(w http.ResponseWriter, r *http.Request) {
	log.Print("List")
	params := []URLShort{
		{Name: "First", URL: "https://example.com/", URLShort: "pgre1"},
		{Name: "First", URL: "https://www.youtube.com/", URLShort: "askf"},
		{Name: "First", URL: "https://github.com/", URLShort: "spsokg"},
	}

	parsed, err := template.ParseFS(templates, "list.html")
	if err != nil {
		log.Print("ParseFS Failed", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	parsed.Execute(w, listParams{params})
}

func Start() {
	log.Print("Start run")
	http.HandleFunc("/", dashboard)
	http.HandleFunc("/list", list)

	log.Print("Listen and Serve")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("HTTP Error - ", err)
	}
}

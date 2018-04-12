package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := populateTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/goapp", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "text/plain")
		writer.Write([]byte("Is this my first Go webapp?"))
	})

	http.HandleFunc("/goappyes", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "text/plain")
		writer.Write([]byte("Yes, this my first Go webapp!"))
	})

	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/img/", fileServer)
	http.Handle("/css/", fileServer)

	http.ListenAndServe(":8080", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}

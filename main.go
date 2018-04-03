package main

import "net/http"

func main() {
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
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	http.ListenAndServe(":8080", nil)
}

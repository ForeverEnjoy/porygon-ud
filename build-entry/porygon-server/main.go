package main

import (
	"flag"
	"net/http"

	service_file "github.com/ForeverEnjoy/porygon-ud/service/file"
)

func main() {
	flag.Parse()

	service_file.Init()

	http.HandleFunc("/file", service_file.Post)

	fileServer := http.FileServer(http.Dir(service_file.GetFileBasePath()))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", 301)
	})

	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		panic(err)
	}
}

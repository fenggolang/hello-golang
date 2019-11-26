package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type ContentData struct {
	Content string
}

var (
	port = flag.String("port", "8080", "-p 8080")
)

func main() {
	flag.Parse()
	dir, _ := os.Getwd()
	var tmpl *template.Template
	if strings.Contains(dir, "hello-golang") {
		tmpl = template.Must(template.ParseFiles("./index.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./hello-golang/index.html"))
	}
	http.HandleFunc("/health", health)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ContentData{
			Content: "Hello Golang " + *port,
		}
		tmpl.Execute(w, data)
	})

	fmt.Printf("http server start success!")
	http.ListenAndServe(":"+*port, nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

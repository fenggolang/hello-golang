package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type ContentData struct {
	Content string
}

func main() {
	dir, _ := os.Getwd()
	var tmpl *template.Template
	if strings.Contains(dir, "hello-golang") {
		tmpl = template.Must(template.ParseFiles("./index.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./hello-golang/index.html"))
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ContentData{
			Content: "Hello Golang",
		}
		tmpl.Execute(w, data)
	})

	fmt.Printf("http server start success!")
	http.ListenAndServe(":8080", nil)
}

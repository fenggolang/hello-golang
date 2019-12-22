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
	if strings.Contains(dir, "hello-golang") { // 本地
		tmpl = template.Must(template.ParseFiles("./index.html"))
	} else if strings.Contains(dir, "app-root") { // golang源码发布，默认路径是/opt/app-root/src
		tmpl = template.Must(template.ParseFiles("./hello-golang/index.html"))
	} else { // golang介质包发布，默认路径是/opt
		tmpl = template.Must(template.ParseFiles("./index.html"))
	}
	http.HandleFunc("/health", health)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/api/v1/version", version)
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
	w.Write([]byte("ok"))
}
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("4.10"))
}

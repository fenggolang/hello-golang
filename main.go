package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type ContentData struct {
	Content string
}

func main() {
	pwd, _ := os.Getwd()
	fmt.Printf("当前路径:%v\n", pwd)
	tmpl := template.Must(template.ParseFiles("./index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ContentData{
			Content: "Hello Golang",
		}
		tmpl.Execute(w, data)
	})

	fmt.Printf("http server start success!")
	http.ListenAndServe(":8080", nil)
}

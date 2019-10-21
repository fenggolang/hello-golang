package main


import (
	"fmt"
	"html/template"
	"net/http"
)


type ContentData struct {
	Content string
}


func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ContentData{
			Content: "Hello Golang",
			}
		tmpl.Execute(w, data)
	})

	fmt.Printf("http server start success!")
	http.ListenAndServe(":8080", nil)
}
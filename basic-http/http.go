package main

import (
	"html/template"
	"log"
	"net/http"
)

type IndexData struct {
	Title   string
	Content string
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	template := template.Must(template.ParseFiles("./index.html"))
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	template.Execute(w, data)
}

func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/index", test)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

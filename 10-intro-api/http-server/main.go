package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	ID      int
	Title   string
	Content string
}

var data = []Article{
	{1, "Title 1", "Content 1"},
	{2, "Title 2", "Content 2"},
	{3, "Title 3", "Content 3"},
	{4, "Title 4", "Content 4"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var datamap = map[string]any{
			"status":  "success",
			"message": "POST success",
		}
		var result, err = json.Marshal(datamap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "PUT" {

	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/books", articles)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}

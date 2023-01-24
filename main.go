package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type post struct {
	UserID int    `json:"userID"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var form = `
<h1>Post #{{.ID}}</h1>
<div>{{printf "The User %d" .UserID}}</div>
<div>{{printf "Title is %s" .Title}}</div>
<div>{{printf "Body is  %s" .Body}}</div>`

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const baseUrl = "https://jsonplaceholder.typicode.com/"
	resp, err := http.Get(baseUrl + r.URL.Path[1:])

	if err != nil {
		//signal an error with http.Error
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	//ensure to close the body that was open else server runs out of sockets
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		//signal an error with http.Error
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	var item post

	err = json.Unmarshal(body, &item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	_template := template.New("sample")
	_template.Parse(form)
	_template.Execute(w, item)
}

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var index *template.Template

func serveIndex(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	log.Println(fmt.Sprintf("%s - %s %s - %s", ip, r.Method, r.URL.Path, r.UserAgent()))

	w.WriteHeader(200)
	index.Execute(w, nil)
}

func serveRobots(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User-agent: *\nAllow: /"))
}

func main() {
	templateStr, _ := ioutil.ReadFile("index.html")
	index, _ = template.New("index").Parse(string(templateStr))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/robots.txt", serveRobots)
	log.Fatal(http.ListenAndServe(":80", nil))
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var res *template.Template

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", 405)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "not found", 404)
		return
	}
	file, err := template.ParseFiles("new.html")
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
	res.Execute(w, file)

}
func handleAsc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", 405)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "400 bad request", 400)
		return
	}
	banner := r.FormValue("banner")
	if banner == " " {
		http.Error(w, "400 bad request", 400)
		return
	}

	call, err := asc(text, banner)
	if err != nil {
		http.Error(w, "500 internal sever error", 500)
		return
	}
	res.Execute(w, call)
}

func asc(text, banner string) (string, error) {
	rawbanner, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return "", err
	}

	replace := strings.ReplaceAll(string(rawbanner), "\r\n", "\n")
	treatedbanner := strings.Split(replace, "\n")
	treatedtext := strings.Split(text, "\n")

	var s strings.Builder

	for _, words := range treatedtext {
		if words == " " {
			s.WriteString("\n")
		}
		for i := 0; i < 8; i++ {
			for _, asci := range words {
				cha := int(9*(asci-32)) + i
				s.WriteString(treatedbanner[cha+1])

			}
		}
	}
	return s.String(), err
}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/handleGet", handleGet)
	fmt.Println("server listening on :8080")
	http.ListenAndServe(":8080", nil)

}

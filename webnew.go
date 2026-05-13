package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func web(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Jesus loves you")
	var err error
	var new *template.Template
	if r.URL.Path != "/w" {
		http.Error(w, "404 Page not found ", http.StatusNotFound)
		return
	}
	new, err = template.ParseFiles("paul.html")
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}
	got := r.FormValue("text")
	banner := r.FormValue("banner")
	str, err := isAsci(got, banner)
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}

	if got == "" {
		// fmt.Fprint(w, "Text empty")
		http.Error(w, "400 Bad Request!!! Enter a text ", http.StatusBadRequest)
		return

	}
	err = new.Execute(w, str)
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}
}

func isAsci(tex, banner string) (string, error) {
	text, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Printf("Error reading banner file, Error: %d", http.StatusNotFound)
		// http.Error(w, "Error: ", http.StatusInternalServerError)
		return "", err
	}
	p := strings.Split(tex, "\\n")
	t := strings.ReplaceAll(string(text), "\r\n", "\n")
	f := strings.Split(t, "\n")

	var s strings.Builder
	for _, words := range p {
		if string(words) == "" {
			s.WriteString("\n")
		}
		for i := 0; i < 8; i++ {
			for _, g := range words {
				ch := int(9*(g-32) + 1)
				s.WriteString(f[ch+i])
			}
			s.WriteString("\n")
		}

	}
	return s.String(), nil

}
func main() {
	http.HandleFunc("/w", web)
	// http.HandleFunc("paul.html", nil)
	fmt.Println("good to go with this link http://localhost:8080/w")
	http.ListenAndServe(":8080", nil)
}

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
	if r.URL.Path != "/" {
		http.Error(w, "404 Page not found ", http.StatusNotFound)
		return
	}
	new, err = template.ParseFiles("paul.html")
	if err != nil {
		http.Error(w, "500 Internal server error1", http.StatusInternalServerError)
		return
	}

	new.Execute(w, "")
	// if err != nil {
	// 	http.Error(w, "500 Internal server error3", http.StatusInternalServerError)
	// 	return
	// }
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var new *template.Template
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var str string
	got := r.FormValue("text")
	if got == "" {
		// fmt.Fprint(w, "Text empty")
		http.Error(w, "400 Bad Request!!! Enter a text ", http.StatusBadRequest)
		return

	}
	banner := r.FormValue("banner")
	if banner == "" {
		http.Error(w, "400 Bad Request !!!", http.StatusBadRequest)
	}
	str, err := isAsci(got, banner)
	if err != nil {
		http.Error(w, "500 Internal server error2", http.StatusInternalServerError)
		return
	}
	new, err = template.ParseFiles("paul.html")
	if err != nil {
		http.Error(w, "500 Internal server error1", http.StatusInternalServerError)
		return
	}

	new.Execute(w, str)

}

func isAsci(rawText, banner string) (string, error) {

	rawBanner, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Printf("Error reading banner file, Error: %d", http.StatusInternalServerError)
		// http.Error(w, "Error: ", http.StatusInternalServerError)
		return "", err
	}

	treatedBanner := strings.Split(strings.ReplaceAll(string(rawBanner), "\r\n", "\n"), "\n")

	temp := strings.ReplaceAll(rawText, "\r\n", "\n")
	temp1 := strings.ReplaceAll(temp, "\\n", "\n")
	temp2 := strings.ReplaceAll(temp1, "\\r", "")
	treatedWord := strings.Split(temp2, "\n")

	// p := strings.Split(tex, "\r\n")
	// t := strings.ReplaceAll(string(text), "\r\n", "\n")
	// f := strings.Split(t, "\n")

	var s strings.Builder
	for _, words := range treatedWord {
		if string(words) == "" {
			s.WriteString("\n")
		}
		for i := 0; i < 8; i++ {
			for _, g := range words {
				ch := int(9*(g-32) + 1)
				s.WriteString(treatedBanner[ch+i])
			}
			s.WriteString("\n")
		}

	}
	return s.String(), nil

}
func main() {
	d := http.Dir(".")
	h := http.FileServer(d)
	http.Handle("/style.css", h)
	http.HandleFunc("/", web)
	http.HandleFunc("/ascii", postHandler)
	fmt.Println("good to go with this link http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

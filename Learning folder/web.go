package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

// func isWeb(t http.ResponseWriter, j *http.Request) {
// 	// tem, _ := template.ParseFiles("index.html")
// 	// tem.Execute(t, nil)
// 	// fmt.Fprintf(t, "My name is Paul i am the founder of greatmercy fashion house!")
// }

// func wel(t http.ResponseWriter, j *http.Request) {
// 	fmt.Fprint(t, "JESUS loves you")
// }

tem *template.Template
// func isAsci(w http.ResponseWriter, r *http.Request) {
// 	tex := r.FormValue("string")
// 	out := fmt.Sprintf("you typed %s", tex)
// 	// now.Execute(w, out)
// 	fmt.Fprint(w, out)
// 	banner := r.FormValue("banner")
// 	text, err := os.ReadFile(banner + ".txt")
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	p := strings.Split(tex, "\\n")
// 	t := strings.ReplaceAll(string(text), "\r\n", "\n")
// 	f := strings.Split(t, "\n")
	
// 	var s strings.Builder
// 	for _, words := range p {
// 		if string(words) == "" {
// 			s.WriteString("\n")
// 		}
// 		for i := 0; i < 8; i++ {
// 			for _, g := range words {
// 				ch := int(9*(g-32) + 1)
// 				s.WriteString(f[ch+i])
// 			}
// 			s.WriteString("\n")
// 		}
		
// 	}
// 	tem.Execute(w, s)
	
// }

func main() {
	tem, _ := template.ParseFiles("index.html")
	// http.HandleFunc("/", isWeb)
	http.HandleFunc("/isAsci", isAsci)
	http.Handle("/style.css", http.FileServer(http.Dir(".")))
	fmt.Println("my server has started")
	http.ListenAndServe(":8080", nil)
	// strings.Contains()
}

package main

import (
	"htmlk/template"
	"net/http"
)

var tpl *template.Template

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "page not found", 400)
		return
	}

	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal server error", 500)
	}

	// names := []string{"John", "Paul", "Samuel"}
	// ages := []int{10, 15, 20}
	// var age int
	// fmt.Print("Enter age ")
	// fmt.Scanln(&age)
	// for i := range len(ages) {
	// 	ages[i] = ages[i] + age
	// }
	// fmt.Printf("ages,%v\n", ages)

	// fmt.Println(names)
	// fmt.Println(ages)

}

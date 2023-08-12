package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Name     string
	Workload int
}

type Courses []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Courses{
			{"Go", 40},
			{"Java", 20},
			{"Python", 10},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8383", nil)
}

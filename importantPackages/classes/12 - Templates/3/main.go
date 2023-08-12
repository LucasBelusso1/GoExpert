package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Courses{
		{"GO", 50},
		{"Java", 10},
		{"Python", 30},
	})

	if err != nil {
		panic(err)
	}
}

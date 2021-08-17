package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

type Sage struct {
	Name string
	Rank string
}

func main() {
	sage := Sage{
		Name: "Andrew Hojo",
		Rank: "Most Powerful",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", sage)
	if err != nil {
		log.Fatalln(err)
	}
}

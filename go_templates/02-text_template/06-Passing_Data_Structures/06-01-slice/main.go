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

func main() {
	sage := []string{
		"Andrew",
		"Kairi",
		"Lindy",
		"Obiwan",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", sage)
	if err != nil {
		log.Fatalln(err)
	}
}

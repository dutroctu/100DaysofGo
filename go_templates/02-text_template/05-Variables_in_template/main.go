package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", "Tan Le")
	if err != nil {
		log.Fatalln(err)
	}
}

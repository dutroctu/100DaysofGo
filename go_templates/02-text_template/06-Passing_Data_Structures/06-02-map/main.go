package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	sage := map[string]string{
		"Most Powerful":        "Andrew Hojo",
		"Second Most Powerful": "Kairi Hojo",
		"Only has magic":       "Lindy",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", sage)
	if err != nil {
		log.Fatalln(err)
	}
}

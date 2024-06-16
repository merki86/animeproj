package views

import "html/template"

var Template *template.Template

func init() {
	Template = template.Must(template.ParseGlob("views/*.gohtml"))
}

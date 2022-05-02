package view

import (
	"001_go_env/crawler/frontend/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T){
	template := template.Must(
		template.ParseFiles("template.html"),
		)

	page := model.SearchTest{}
	err := template.Execute(os.Stdout, page)

	if err != nil{
		panic(err)
	}
}

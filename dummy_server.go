package main

import (
	"fmt"
	"html/template"
	"hub/framework"
	"net/http"
)

type LogicDesignerChooserModel struct {
	ErrorMessage string
}

func composeWeb(templateName string, data interface{}, w http.ResponseWriter) {
	t, err := template.ParseFiles(
		framework.GetHtmlTemplatePath(templateName))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		t.Execute(w, data)
	}
}

func handlePostData(data *LogicDesignerChooserModel) {
	data.ErrorMessage = "POST received"
}

func handleUnsupportedHttpMethod(method string, w http.ResponseWriter) {
	fmt.Fprintf(w, "nope")
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
	data := &LogicDesignerChooserModel{""}
	switch r.Method {
	case "GET":
		break
	case "POST":
		handlePostData(data)
	default:
		handleUnsupportedHttpMethod(r.Method, w)
	}
	composeWeb("logic_designer_chooser.html", data, w)
}

func main() {
	http.HandleFunc("/logic_designer", helloWeb)
	http.ListenAndServe(":8080", nil)
}

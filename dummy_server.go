package main

import (
	"fmt"
	"html/template"
	"hub/framework"
	"hub/persistent_storage"
	"net/http"
)

type LogicDesignerChooserModel struct {
	ErrorMessage string
	RuleFiles    []string
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

func handleCommonData(data *LogicDesignerChooserModel) (err error) {
	data.RuleFiles, err = persistent_storage.FilterFiles(
		framework.GetUserDataPath(""),
		persistent_storage.GetExtensionFileInfoMatcher(
			persistent_storage.RuleSetExtension))

	return
}

func handleGetData(data *LogicDesignerChooserModel) (err error) {
	err = handleCommonData(data)
	return
}

func handlePostData(data *LogicDesignerChooserModel) (err error) {
	err = handleCommonData(data)
	data.ErrorMessage = "POST received"
	return
}

func handleUnsupportedHttpMethod(method string, w http.ResponseWriter) {
	fmt.Fprintf(w, "nope")
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
	data := &LogicDesignerChooserModel{"", []string{}}
	var dataHandlingError error
	switch r.Method {
	case "GET":
		dataHandlingError = handleGetData(data)
	case "POST":
		dataHandlingError = handlePostData(data)
	default:
		handleUnsupportedHttpMethod(r.Method, w)
	}

	if dataHandlingError == nil {
		composeWeb("logic_designer_chooser.html", data, w)
	} else {
		fmt.Fprintf(w, "very nope")
	}
}

func main() {
	http.HandleFunc("/logic_designer", helloWeb)
	http.ListenAndServe(":8080", nil)
}

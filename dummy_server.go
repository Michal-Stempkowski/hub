package main

import (
	"fmt"
	"html/template"
	"hub/framework"
	"hub/persistent_storage"
	"hub/server"
	"net/http"
)

type LogicDesignerChooserModel struct {
	ErrorMessage string
	RuleFiles    []string
}

func NewLogicDesignerChooserModel() *LogicDesignerChooserModel {
	return &LogicDesignerChooserModel{"", []string{}}
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

func prepareRulesetList(data *LogicDesignerChooserModel) (err error) {
	data.RuleFiles, err = persistent_storage.FilterFiles(
		framework.GetUserDataPath(""),
		persistent_storage.GetExtensionFileInfoMatcher(
			persistent_storage.RuleSetExtension))

	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (accepted bool) {
	accepted = true

	data := NewLogicDesignerChooserModel()
	err := prepareRulesetList(data)

	renderLogicDesignerMainView(data, err, w)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (accepted bool) {
	accepted = true

	data := NewLogicDesignerChooserModel()
	data.ErrorMessage = "POST received"
	err := prepareRulesetList(data)

	renderLogicDesignerMainView(data, err, w)
	return
}

func renderLogicDesignerMainView(
	data *LogicDesignerChooserModel, dataHandlingError error, w http.ResponseWriter) {
	if dataHandlingError == nil {
		composeWeb("logic_designer_chooser.html", data, w)
	} else {
		fmt.Fprintf(w, "very nope\n")
		fmt.Fprintf(w, dataHandlingError.Error())
	}
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
	handler := server.NewRequestBroker(
		map[string][]server.RequestHandler{
			"GET": []server.RequestHandler{
				handleGet,
			},
			"POST": []server.RequestHandler{
				handlePost,
			},
		})

	handler.Resolve(w, r)
}

func main() {
	r := &server.RequestBroker{}
	http.HandleFunc("/logic_designer", helloWeb)
	http.HandleFunc("/logic_designer2", r.Resolve)
	http.ListenAndServe(":8080", nil)
}

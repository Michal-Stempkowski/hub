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

func handleGetData(data *LogicDesignerChooserModel) (err error) {
	err = prepareRulesetList(data)
	return
}

func handlePostData(data *LogicDesignerChooserModel) (err error) {
	err = prepareRulesetList(data)
	data.ErrorMessage = "POST received"
	return
}

func handleUnsupportedHttpMethod(method string, w http.ResponseWriter) error {
	fmt.Fprintf(w, "nope\n")
	return fmt.Errorf("This http method is unsupported %s", method)
}

func logicDesignerFlow(
	data_filler func(*LogicDesignerChooserModel) error) server.RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) bool {
		data := &LogicDesignerChooserModel{"", []string{}}
		dataHandlingError := data_filler(data)

		if dataHandlingError == nil {
			composeWeb("logic_designer_chooser.html", data, w)
		} else {
			fmt.Fprintf(w, "very nope\n")
			fmt.Fprintf(w, dataHandlingError.Error())
		}
		return true
	}
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
	handler := server.NewRequestBroker(
		map[string][]server.RequestHandler{
			"GET": []server.RequestHandler{
				logicDesignerFlow(handleGetData),
			},
			"POST": []server.RequestHandler{
				logicDesignerFlow(handlePostData),
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

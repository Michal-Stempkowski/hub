package server

import (
	"fmt"
	"html/template"
	"hub/framework"
	"net/http"
	"strings"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

func StartServer(urlHandlerMap map[string]HttpHandler) {
	for url, handler := range urlHandlerMap {
		http.HandleFunc(url, handler)
	}

	http.ListenAndServe(":8080", nil)
}

func MergeHandlerMaps(
	maps ...map[string]HttpHandler) map[string]HttpHandler {

	initSize := 0
	for _, m := range maps {
		initSize += len(m)
	}

	result := make(map[string]HttpHandler, initSize)
	for _, m := range maps {
		for k, v := range m {
			if _, alreadyHasKey := result[k]; alreadyHasKey {
				panic("Multiple handlers under same URL path: " + k)
			}
			result[k] = v
		}
	}

	return result
}

func RequestActionMatches(r *http.Request, action string) bool {
	return strings.HasPrefix(r.FormValue("action"), action+":")
}

func RenderSimpleView(
	data interface{},
	template string,
	dataHandlingError error,
	w http.ResponseWriter) {

	if dataHandlingError == nil {
		composeWeb(template, data, w)
	} else {
		simpleHandleError(dataHandlingError, w)
	}
}

func ParsePostData(w http.ResponseWriter, r *http.Request) (accepted bool) {
	err := r.ParseForm()
	if err != nil {
		accepted = true
		simpleHandleError(err, w)
	}

	return
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

func simpleHandleError(err error, w http.ResponseWriter) {
	fmt.Fprintf(w, "very nope\n")
	fmt.Fprintf(w, err.Error())
}

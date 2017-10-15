package logic_designer

import (
	"hub/framework"
	"hub/persistent_storage"
	"hub/server"
	"net/http"
	"strings"
)

func GetHandlers() map[string]server.HttpHandler {
	return map[string]server.HttpHandler{
		"/logic_designer":        logicDesigner,
		"/logic_designer/rename": logicDesignerRename,
	}
}

type logicDesignerChooserModel struct {
	ErrorMessage string
	RuleFiles    []string
}

func newLogicDesignerChooserModel() *logicDesignerChooserModel {
	return &logicDesignerChooserModel{"", []string{}}
}

func prepareRulesetList(data *logicDesignerChooserModel) (err error) {
	data.RuleFiles, err = persistent_storage.FilterFiles(
		framework.GetUserDataPath(""),
		persistent_storage.GetExtensionFileInfoMatcher(
			persistent_storage.RuleSetExtension))

	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (accepted bool) {
	accepted = true

	data := newLogicDesignerChooserModel()
	err := prepareRulesetList(data)

	server.RenderSimpleView(
		data, "logic_designer_chooser.html", err, w)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (accepted bool) {
	accepted = true

	data := newLogicDesignerChooserModel()
	data.ErrorMessage = "POST received"
	err := prepareRulesetList(data)

	server.RenderSimpleView(
		data, "logic_designer_chooser.html", err, w)
	return
}

func handleRename(w http.ResponseWriter, r *http.Request) bool {
	if !strings.HasPrefix(r.FormValue("action"), "RequestRename:") {
		return false
	}

	url := strings.Join([]string{r.URL.Path, "rename"}, "/")

	http.Redirect(w, r, url, http.StatusPermanentRedirect)

	return true
}

func logicDesigner(w http.ResponseWriter, r *http.Request) {
	handler := server.NewRequestBroker(
		map[string][]server.RequestHandler{
			"GET": []server.RequestHandler{
				handleGet,
			},
			"POST": []server.RequestHandler{
				server.ParsePostData,
				handleRename,
				handlePost,
			},
		})

	handler.Resolve(w, r)
}

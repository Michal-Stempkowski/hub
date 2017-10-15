package logic_designer

import (
	"fmt"
	"hub/framework"
	"hub/persistent_storage"
	"hub/persistent_storage/file_operation_status"
	"hub/server"
	"net/http"
	"strings"
)

const (
	RequestRename  = "RequestRename"
	VetoableRename = "VetoableRename"
	Cancel         = "Cancel"
)

func logicDesignerRename(w http.ResponseWriter, r *http.Request) {
	handler := server.NewRequestBroker(
		map[string][]server.RequestHandler{
			"POST": []server.RequestHandler{
				handleRequestRename,
				handleCancel,
				handleVetoableRename,
			},
		})

	handler.Resolve(w, r)
}

func handleRequestRename(w http.ResponseWriter, r *http.Request) bool {
	if !server.RequestActionMatches(r, RequestRename) {
		return false
	}

	oldName := getOldNameFromRequest(r)

	data := defaultRenameModel(oldName)

	server.RenderSimpleView(data, "ok_cancel_text_question.html", nil, w)

	return true
}

func handleVetoableRename(w http.ResponseWriter, r *http.Request) bool {
	if !server.RequestActionMatches(r, VetoableRename) {
		return false
	}

	oldName := getOldNameFromRequest(r)
	newName := r.FormValue("value")

	var err error
	if newName == "" {
		err = fmt.Errorf(file_operation_status.EmptyName)
	} else {
		err = persistent_storage.SafeRenameForExtension(
			framework.GetUserDataPath(oldName),
			framework.GetUserDataPath(r.FormValue("value")),
			persistent_storage.RuleSetExtension)
	}

	if err == nil {
		http.Redirect(w, r, "/logic_designer", http.StatusSeeOther)
	} else {
		data := defaultRenameModel(oldName)
		data.ErrorMessage = err.Error()

		server.RenderSimpleView(data, "ok_cancel_text_question.html", nil, w)
	}

	return true
}

func handleCancel(w http.ResponseWriter, r *http.Request) bool {
	if !server.RequestActionMatches(r, Cancel) {
		return false
	}

	http.Redirect(w, r, "/logic_designer", http.StatusSeeOther)

	return true
}

func defaultRenameModel(oldName string) *server.OkCancelTextQuestionModel {
	return &server.OkCancelTextQuestionModel{
		Title: fmt.Sprintf(
			"CharacterHub - LogicDesigner - renaming %v", oldName),
		Question:     fmt.Sprintf("What should be new name for %v?", oldName),
		ErrorMessage: "",
		Destination:  "/logic_designer/rename",
		Ok:           VetoableRename + ":" + oldName,
		OkLabel:      "Rename",
		Cancel:       Cancel + ":" + oldName,
		CancelLabel:  "Cancel"}
}

func getOldNameFromRequest(r *http.Request) string {
	actionData := strings.Split(r.FormValue("action"), ":")
	return actionData[len(actionData)-1]
}

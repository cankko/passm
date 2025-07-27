package handler

import (
	"fmt"
	"net/http"

	"passm/internal/helper"
	"passm/internal/model"
	"passm/internal/repository"
	"passm/internal/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != helper.ListPath {
		helper.TemplateNotFound(w)
		return
	}

	entries, err := repository.GetEntries()
	if err != nil {
		helper.TemplateInternalError(fmt.Sprintf("Error reading entries: %s", err), w)
		return
	}

	var modifyEntries []model.Entry

	for _, item := range entries {
		modifyEntries = append(
			modifyEntries,
			model.Entry{
				ID:       item.ID,
				Source:   item.Source,
				User:     item.User,
				Password: service.Decrypt(item.Password),
			})
	}

	helper.TemplateParse("home.html", modifyEntries, w)
}

package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"passm/internal/helper"
	"passm/internal/repository"
	"passm/internal/service"
)

func CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	if "" != r.FormValue("password") {
		encryptPassword := service.Encrypt(r.FormValue("password"))
		err = repository.CreateEntry(r.FormValue("source"), r.FormValue("user"), encryptPassword)
		if err != nil {
			helper.TemplateInternalError(fmt.Sprintf("Error creating entry: %s", err), w)
			return
		}
	}

	http.Redirect(w, r, helper.ListPath, http.StatusSeeOther)
}

func DeleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	err := repository.DeleteEntry(id)
	if err != nil {
		helper.TemplateInternalError(fmt.Sprintf("Error deleting entry: %s", err), w)
		return
	}

	http.Redirect(w, r, helper.ListPath, http.StatusSeeOther)
}

func UpdateEntryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(r.FormValue("id"))

	encryptPassword := service.Encrypt(r.FormValue("password"))
	err = repository.UpdateEntry(id, r.FormValue("source"), r.FormValue("user"), encryptPassword)
	if err != nil {
		helper.TemplateInternalError(fmt.Sprintf("Error updating entry: %s", err), w)
		return
	}

	http.Redirect(w, r, helper.ListPath, http.StatusSeeOther)
}

package handler

import (
	"fmt"
	"log"
	"net/http"

	"passm/internal/helper"
	"passm/internal/repository"
	"passm/internal/service"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := service.GetFlashMessage(r, w)
	helper.TemplateParse("login.html", data, w)
}

func LoginTryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	getUsers, err := repository.LoadUsers()
	if err != nil {
		fmt.Println("Error reading items:", err)
	}

	for _, getUser := range getUsers {
		if nil == bcrypt.CompareHashAndPassword([]byte(getUser.MainPassword), []byte(r.FormValue("passname"))) {
			service.CreateSession(w, r, getUser.ID)
			http.Redirect(w, r, helper.ListPath, http.StatusSeeOther)
			return
		}
	}

	service.SetFlashMessage("wrong password!", r, w)

	http.Redirect(w, r, helper.LoginFormPath, http.StatusSeeOther)
	return
}

func LoginLogout(w http.ResponseWriter, r *http.Request) {
	service.DestroySession(w, r)
	http.Redirect(w, r, helper.LoginFormPath, http.StatusSeeOther)
}

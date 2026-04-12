package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"passm/internal/db"
	"passm/internal/handler"
	"passm/internal/helper"
	"passm/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/static"))
	router.Handle(http.MethodGet+" /static/", http.StripPrefix("/static", fs))
	router.Handle(http.MethodGet+" /robots.txt", fs)

	router.HandleFunc(http.MethodGet+" /", handler.HomeHandler)
	router.HandleFunc(http.MethodGet+" /login/form/", handler.LoginHandler)
	router.HandleFunc(http.MethodPost+" /login/try/", handler.LoginTryHandler)
	router.HandleFunc(http.MethodGet+" /logout/", handler.LoginLogout)
	router.HandleFunc(http.MethodPost+" /entry_create/", handler.CreateEntryHandler)
	router.HandleFunc(http.MethodGet+" /entry_delete/{id}/", handler.DeleteEntryHandler)
	router.HandleFunc(http.MethodPost+" /entry_update/", handler.UpdateEntryHandler)

	server := &http.Server{
		Addr:    ":3200",
		Handler: middleware(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isLoginPaths := strings.Contains(r.URL.Path, "/login/")
		isStaticPaths := strings.Contains(r.URL.Path, "/static/")
		isRobotsTxt := strings.Contains(r.URL.Path, "robots.txt")

		if !isStaticPaths {
			fmt.Println("From middleware, static: " + r.URL.Path)
			helper.SetGlobalSession(r)
		}

		if !service.IsLoggedIn() && !isLoginPaths && !isStaticPaths && !isRobotsTxt {
			http.Redirect(w, r, helper.LoginFormPath, http.StatusSeeOther)
			return
		}

		if service.IsLoggedIn() && isLoginPaths {
			fmt.Println("From middleware, redirect to list")
			http.Redirect(w, r, helper.ListPath, http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

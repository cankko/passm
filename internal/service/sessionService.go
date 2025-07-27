package service

import (
	"fmt"
	"net/http"

	"passm/internal/helper"
)

var timeOutInSeconds = 600

func CreateSession(w http.ResponseWriter, r *http.Request, userId int) {
	helper.Session.Options.MaxAge = timeOutInSeconds
	helper.Session.Values["isLoggedIn"] = true
	helper.Session.Values["user_id"] = userId
	err := helper.Session.Save(r, w)
	if err != nil {
		return
	}
}

func IsLoggedIn() bool {
	isLoggedIn := true
	if auth, ok := helper.Session.Values["isLoggedIn"].(bool); !ok || !auth {
		isLoggedIn = false
	}

	return isLoggedIn
}

func DestroySession(w http.ResponseWriter, r *http.Request) {
	helper.Session.Options.MaxAge = -1 // destroy session
	err := helper.Session.Save(r, w)
	if err != nil {
		return
	}
}

func SetFlashMessage(msg string, r *http.Request, w http.ResponseWriter) {
	helper.Session.AddFlash(msg, "msg")
	err := helper.Session.Save(r, w)
	if err != nil {
		fmt.Println("Error adding flash msg:", err)
	}
}

type Message struct {
	Msg string
}

func GetFlashMessage(r *http.Request, w http.ResponseWriter) *Message {
	message := Message{}
	fm := helper.Session.Flashes("msg")
	if fm != nil {
		message = Message{
			Msg: fm[0].(string),
		}
	}
	err := helper.Session.Save(r, w)
	if err != nil {
		fmt.Println("Error getting flash msg:", err)
	}

	return &message
}

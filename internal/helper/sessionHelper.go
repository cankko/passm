package helper

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Session *sessions.Session
var Store = sessions.NewCookieStore([]byte("kRSZlkTElTpWNRp6hphjupBEuEGpLmRV"))

func SetGlobalSession(r *http.Request) {
	session, _ := Store.Get(r, "session-name")
	Session = session
}

package handler

import (
	"entry_task/database"
	"github.com/gorilla/sessions"
)

type Handler struct {
	DB           *database.MyDB
	SessionStore sessions.Store
}

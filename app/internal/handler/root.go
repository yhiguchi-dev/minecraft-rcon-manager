package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

var userGetHandler GetHandler
var userItemPostHandler PostHandler
var whitelistUserGetHandler GetHandler
var whitelistPostHandler PostHandler
var whitelistUserPostHandler PostHandler
var whitelistUserDeleteHandler DeleteHandler

func NewRootHandler(_userGetHandler GetHandler, _userItemPostHandler PostHandler, _whitelistUserGetHandler GetHandler, _whitelistPostHandler PostHandler, _whitelistUserPostHandler PostHandler, _whitelistUserDeleteHandler DeleteHandler) {
	userGetHandler = _userGetHandler
	userItemPostHandler = _userItemPostHandler
	whitelistUserGetHandler = _whitelistUserGetHandler
	whitelistPostHandler = _whitelistPostHandler
	whitelistUserPostHandler = _whitelistUserPostHandler
	whitelistUserDeleteHandler = _whitelistUserDeleteHandler
}

type GetHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type PostHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
}

type DeleteHandler interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := fmt.Fprint(w, "Method not allowed.\n")
		if err != nil {
			return
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if url == "/users" {
		userGetHandler.Get(w, r)
	} else if url == "/whitelist/users" {
		whitelistUserGetHandler.Get(w, r)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	userItemRegex := regexp.MustCompile(`/users/[a-zA-Z0-9_]+/item`)
	whitelistRegex := regexp.MustCompile(`/whitelist/users/[a-zA-Z0-9_]+$`)
	url := r.URL.Path
	if userItemRegex.MatchString(url) {
		userItemPostHandler.Post(w, r)
	} else if url == "/whitelist" {
		whitelistPostHandler.Post(w, r)
	} else if whitelistRegex.MatchString(url) {
		whitelistUserPostHandler.Post(w, r)
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	whitelistRegex := regexp.MustCompile(`/whitelist/users/[a-zA-Z0-9_]+$`)
	url := r.URL.Path
	if whitelistRegex.MatchString(url) {
		whitelistUserDeleteHandler.Delete(w, r)
	}
}

package whitelist

import (
	"app/internal/handler"
	"app/internal/rcon"
	"fmt"
	"net/http"
	"path/filepath"
)

type whitelistUserPostHandler struct {
	operator rcon.Operator
}

func NewWhitelistUserPostHandler(operator rcon.Operator) handler.PostHandler {
	return &whitelistUserPostHandler{operator}
}

func (handler whitelistUserPostHandler) Post(w http.ResponseWriter, r *http.Request) {
	_, userId := filepath.Split(r.URL.Path)
	whitelist, err := handler.operator.AddUserToWhitelist(userId)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, whitelist)
	if err != nil {
		return
	}
}

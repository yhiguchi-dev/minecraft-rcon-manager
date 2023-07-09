package whitelist

import (
	"app/internal/handler"
	"app/internal/rcon"
	"fmt"
	"net/http"
	"path/filepath"
)

type whitelistUserDeleteHandler struct {
	operator rcon.Operator
}

func NewWhitelistUserDeleteHandler(operator rcon.Operator) handler.DeleteHandler {
	return &whitelistUserDeleteHandler{operator}
}

func (handler whitelistUserDeleteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_, userId := filepath.Split(r.URL.Path)
	whitelist, err := handler.operator.RemoveUserToWhitelist(userId)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, whitelist)
	if err != nil {
		return
	}
}

package whitelist

import (
	"app/internal/handler"
	"app/internal/rcon"
	"encoding/json"
	"fmt"
	"net/http"
)

type whitelistPostHandler struct {
	operator rcon.Operator
}

func NewWhitelistPostHandler(operator rcon.Operator) handler.PostHandler {
	return &whitelistPostHandler{operator}
}

func (handler whitelistPostHandler) Post(w http.ResponseWriter, r *http.Request) {
	var request = whitelistPostRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return
	}
	if request.Enabled == true {
		whitelist, err := handler.operator.ActivateWhitelist()
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(w, whitelist)
		if err != nil {
			return
		}
	} else {
		whitelist, err := handler.operator.DeactivateWhitelist()
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(w, whitelist)
		if err != nil {
			return
		}
	}
}

type whitelistPostRequest struct {
	Enabled bool `json:"enabled"`
}

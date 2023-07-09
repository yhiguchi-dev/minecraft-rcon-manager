package whitelist

import (
	"app/internal/handler"
	"app/internal/rcon"
	"encoding/json"
	"net/http"
)

type whitelistGetHandler struct {
	operator rcon.Operator
}

func NewWhitelistGetHandler(operator rcon.Operator) handler.GetHandler {
	return &whitelistGetHandler{operator}
}

func (handler *whitelistGetHandler) Get(w http.ResponseWriter, _ *http.Request) {
	list, err := handler.operator.GetWhitelist()
	if err != nil {
		return
	}
	var response = whitelistGetResponse{Users: list}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

type whitelistGetResponse struct {
	Users []string `json:"users"`
}

package user

import (
	"app/internal/handler"
	"app/internal/rcon"
	"encoding/json"
	"net/http"
)

type userGetHandler struct {
	operator rcon.Operator
}

func NewUserGetHandler(operator rcon.Operator) handler.GetHandler {
	return &userGetHandler{operator}
}

func (handler userGetHandler) Get(w http.ResponseWriter, _ *http.Request) {
	list, err := handler.operator.GetUserList()
	if err != nil {
		return
	}
	var response = userGetResponse{Users: list}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

type userGetResponse struct {
	Users []string `json:"users"`
}

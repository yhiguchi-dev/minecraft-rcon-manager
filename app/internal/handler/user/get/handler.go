package get

import (
	"app/internal/rcon"
	"encoding/json"
	"net/http"
)

var operator rcon.Operator

func NewUserListHandler(_operator rcon.Operator) {
	operator = _operator
}

func UserListHandler(w http.ResponseWriter, _ *http.Request) {
	list, err := operator.GetUserList()
	if err != nil {
		return
	}
	var response = userListResponse{Users: list}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

type userListResponse struct {
	Users []string `json:"users"`
}

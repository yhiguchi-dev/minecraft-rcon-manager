package get

import (
	"app/internal/rcon"
	"fmt"
	"net/http"
)

var operator rcon.Operator

func NewUserListHandler(_operator rcon.Operator) {
	operator = _operator
}

func UserListHandler(writer http.ResponseWriter, _ *http.Request) {
	list, err := operator.GetUserList()
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(writer, list)
	if err != nil {
		return
	}
}

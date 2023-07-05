package post

import (
	"fmt"
	"net/http"
	"rcon-client/internal/rcon"
)

var operator rcon.Operator

func NewPostUserListHandler(_operator rcon.Operator) {
	operator = _operator
}

func PostUserListHandler(writer http.ResponseWriter, _ *http.Request) {
	list, err := operator.GetUserList()
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(writer, list)
	if err != nil {
		return
	}
}

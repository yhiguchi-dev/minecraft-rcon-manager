package post

import (
	"app/internal/rcon"
	"fmt"
	"net/http"
	"strings"
)

var operator rcon.Operator

func NewUserItemHandler(_operator rcon.Operator) {
	operator = _operator
}

func UserItemHandler(writer http.ResponseWriter, request *http.Request) {
	pathParam := strings.TrimPrefix(request.URL.Path, "/users/")
	paths := strings.Split(pathParam, "/")
	userId := paths[0]
	itemId := paths[2]
	list, err := operator.GiveItemToUser(userId, itemId)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(writer, list)
	if err != nil {
		return
	}
}

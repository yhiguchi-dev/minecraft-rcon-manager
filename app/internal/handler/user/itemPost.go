package user

import (
	"app/internal/handler"
	"app/internal/rcon"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type userItemPostHandler struct {
	operator rcon.Operator
}

func NewUserItemPostHandler(operator rcon.Operator) handler.PostHandler {
	return &userItemPostHandler{operator}
}

func (handler userItemPostHandler) Post(w http.ResponseWriter, r *http.Request) {
	compile := regexp.MustCompile(`/whitelist/users/([a-zA-Z0-9_]+)/item`)
	group := compile.FindSubmatch([]byte(r.URL.Path))
	userId := string(group[1])
	var request = userItemRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	list, err := handler.operator.GiveItemToUser(userId, request.ItemId, request.Amount)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, list)
	if err != nil {
		return
	}
}

type userItemRequest struct {
	ItemId string `json:"item_id"`
	Amount int    `json:"amount"`
}

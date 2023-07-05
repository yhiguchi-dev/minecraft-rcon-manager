package post

import (
	"app/internal/rcon"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

var operator rcon.Operator

func NewUserItemHandler(_operator rcon.Operator) {
	operator = _operator
}

func UserItemHandler(w http.ResponseWriter, r *http.Request) {
	pathParam := strings.TrimPrefix(r.URL.Path, "/users/")
	_, userId := filepath.Split(pathParam)
	var request struct {
		ItemId string `json:"item_id"`
		Amount int    `json:"amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	list, err := operator.GiveItemToUser(userId, request.ItemId, request.Amount)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, list)
	if err != nil {
		return
	}
}

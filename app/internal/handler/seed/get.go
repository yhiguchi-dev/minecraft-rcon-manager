package seed

import (
	"app/internal/handler"
	"app/internal/rcon"
	"fmt"
	"net/http"
)

type seedGetHandler struct {
	operator rcon.Operator
}

func NewSeedGetHandler(operator rcon.Operator) handler.GetHandler {
	return &seedGetHandler{operator}
}

func (handler seedGetHandler) Get(w http.ResponseWriter, _ *http.Request) {
	response, err := handler.operator.GetSeed()
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, response)
	if err != nil {
		return
	}
}

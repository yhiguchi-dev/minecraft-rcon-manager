package rcon

import "fmt"

type Operator interface {
	GetUserList() (string, error)
	GiveItemToUser(userId string, itemId string) (string, error)
}

type operator struct {
	client Client
}

func NewOperator(client Client) Operator {
	return &operator{client}
}

func (operator *operator) GetUserList() (string, error) {
	packet, err := operator.client.Send("list")
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

func (operator *operator) GiveItemToUser(userId string, itemId string) (string, error) {
	packet, err := operator.client.Send(fmt.Sprintf("give %s %s", userId, itemId))
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

package rcon

import (
	"fmt"
	"strings"
)

type Operator interface {
	GetUserList() ([]string, error)
	GiveItemToUser(userId string, itemId string, amount int) (string, error)
	GetWhitelist() ([]string, error)
	ActivateWhitelist() (string, error)
	DeactivateWhitelist() (string, error)
	AddUserToWhitelist(userId string) (string, error)
	RemoveUserToWhitelist(userId string) (string, error)
	GetSeed() (string, error)
}

type operator struct {
	client Client
}

func NewOperator(client Client) Operator {
	return &operator{client}
}

func (operator *operator) GetUserList() ([]string, error) {
	packet, err := operator.client.Send("list")
	if err != nil {
		return []string{}, err
	}
	payload := string(packet.Payload)
	split := strings.Split(payload, ":")
	_users := strings.Split(split[1], ",")
	users := _users[:]
	for i, v := range users {
		users[i] = strings.TrimSpace(v)
	}
	return users, nil
}

func (operator *operator) GiveItemToUser(userId string, itemId string, amount int) (string, error) {
	packet, err := operator.client.Send(fmt.Sprintf("give %s %s %d", userId, itemId, amount))
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

func (operator *operator) GetWhitelist() ([]string, error) {
	packet, err := operator.client.Send("whitelist list")
	if err != nil {
		return []string{}, err
	}
	payload := string(packet.Payload)
	split := strings.Split(payload, ":")
	_whitelistUsers := strings.Split(split[1], ",")
	whitelistUsers := _whitelistUsers[:]
	for i, v := range whitelistUsers {
		whitelistUsers[i] = strings.TrimSpace(v)
	}
	return whitelistUsers, nil
}

func (operator *operator) ActivateWhitelist() (string, error) {
	packet, err := operator.client.Send("whitelist on")
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

func (operator *operator) DeactivateWhitelist() (string, error) {
	packet, err := operator.client.Send("whitelist off")
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

func (operator *operator) AddUserToWhitelist(userId string) (string, error) {
	packet, err := operator.client.Send(fmt.Sprintf("whitelist add %s", userId))
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

func (operator *operator) RemoveUserToWhitelist(userId string) (string, error) {
	packet, err := operator.client.Send(fmt.Sprintf("whitelist remove %s", userId))
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

func (operator *operator) GetSeed() (string, error) {
	packet, err := operator.client.Send("seed")
	if err != nil {
		return "", err
	}
	payload := string(packet.Payload)
	return payload, nil
}

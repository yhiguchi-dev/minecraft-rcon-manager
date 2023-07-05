package rcon

type Operator interface {
	GetUserList() (string, error)
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

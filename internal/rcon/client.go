package rcon

import (
	"bytes"
	"errors"
	"net"
)

const (
	login           Type = 3
	command              = 2
	commandResponse      = 0
)

type Type int32

type Packet struct {
	Length    int32
	RequestId int32
	Type      Type
	Payload   []byte
	Padding   []byte
}

type Client struct {
	conn net.Conn
	id   int32
}

func NewClient(conn net.Conn, pass string) (Client, error) {
	client := Client{
		conn: conn,
	}
	packet := Packet{
		Type:    login,
		Payload: []byte(pass),
	}
	send, err := client.send(packet)
	if err != nil {
		return Client{}, err
	}
	if send.RequestId == -1 {
		return Client{}, errors.New("incorrect password")
	}
	if send.Type != 2 || send.RequestId != client.id {
		return client, errors.New("illegal state")
	}
	return client, nil
}

func (client Client) Send(exec string) error {
	packet := Packet{
		Type:    command,
		Payload: []byte(exec),
	}
	_, err := client.send(packet)
	if err != nil {
		return err
	}
	return nil
}

func (client Client) send(packet Packet) (Packet, error) {
	err := client.write(packet)
	if err != nil {
		return Packet{}, err
	}
	read, err := client.read()
	if err != nil {
		return read, err
	}
	return read, nil
}

func (client Client) write(packet Packet) error {
	packet.Length = int32(4 + 4 + len(packet.Payload) + 2)
	packet.RequestId = client.id
	packet.Padding = []byte{0x0, 0x0}
	operator := binaryOperator{
		buf: new(bytes.Buffer),
	}
	operator.Write(packet.Length)
	operator.Write(packet.RequestId)
	operator.Write(packet.Type)
	operator.Write(packet.Payload)
	operator.Write(packet.Padding)
	if operator.err != nil {
		return operator.err
	}
	_, err := client.conn.Write(operator.buf.Bytes())
	if err != nil {
		return err
	}
	client.id += 1
	return nil
}

func (client Client) read() (Packet, error) {
	const bufSize = 4096
	b := make([]byte, bufSize)
	read, err := client.conn.Read(b)
	if err != nil {
		return Packet{}, err
	}
	operator := binaryOperator{buf: bytes.NewBuffer(b)}
	packet := Packet{}

	operator.Read(&packet.Length)
	operator.Read(&packet.RequestId)
	operator.Read(&packet.Type)
	body := [bufSize - 12]byte{}
	operator.Read(&body)
	if operator.err != nil {
		return Packet{}, operator.err
	}
	packet.Payload = body[:read-12]
	return packet, nil
}

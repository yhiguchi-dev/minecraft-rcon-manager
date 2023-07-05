package rcon

import (
	"bytes"
	"errors"
	"fmt"
	"log"
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
		id:   0,
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
	if send.Type != command || send.RequestId != client.id {
		return client, errors.New("illegal state")
	}
	return client, nil
}

func (client Client) Send(exec string) (Packet, error) {
	packet := Packet{
		Type:    command,
		Payload: []byte(exec),
	}
	send, err := client.send(packet)
	if err != nil {
		return Packet{}, err
	}
	if send.Type != commandResponse {
		return Packet{}, errors.New("incorrect response")
	}
	return send, nil
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
	packet.RequestId = client.id
	packet.Padding = []byte{0x0, 0x0}
	packet.Length = int32(4 + 4 + len(packet.Payload) + len(packet.Padding))
	operator := binaryOperator{
		buf: new(bytes.Buffer),
	}
	log.Println(fmt.Sprintf("length:%d", packet.Length))
	log.Println(fmt.Sprintf("requestId:%d", packet.RequestId))
	log.Println(fmt.Sprintf("type:%d", packet.Type))
	log.Println(fmt.Sprintf("payload:%s", packet.Payload))
	log.Println(packet.Padding)
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
	packet.Payload = body[:read-14]
	packet.Padding = body[read-14 : read-12]
	log.Println(fmt.Sprintf("read size:%d", read))
	log.Println(fmt.Sprintf("length:%d", packet.Length))
	log.Println(fmt.Sprintf("requestId:%d", packet.RequestId))
	log.Println(fmt.Sprintf("type:%d", packet.Type))
	log.Println(fmt.Sprintf("payload:%s", packet.Payload))
	log.Println(packet.Padding)
	return packet, nil
}

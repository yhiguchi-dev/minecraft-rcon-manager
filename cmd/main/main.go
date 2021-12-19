package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"rcon-client/internal/rcon"
	"time"
)

func main() {
	host := flag.String("host", "", "host name")
	port := flag.Int("port", 0, "port number")
	pass := flag.String("pass", "", "rcon password for minecraft server")
	exec := flag.String("exec", "", "exec command to minecraft server")
	flag.Parse()
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", *host, *port), 3*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	client, err := rcon.NewClient(conn, *pass)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Send(*exec)
	if err != nil {
		log.Fatal(err)
	}
}

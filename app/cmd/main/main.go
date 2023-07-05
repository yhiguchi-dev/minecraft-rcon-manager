package main

import (
	"app/internal/handler/user/get"
	"app/internal/handler/user/post"
	"app/internal/rcon"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	rconServerHost := os.Getenv("RCON_SERVER_HOST")
	rconServerPort := os.Getenv("RCON_SERVER_PORT")
	rconServerPassword := os.Getenv("RCON_SERVER_PASSWORD")

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", rconServerHost, rconServerPort), 3*time.Second)

	if err != nil {
		log.Fatal(err)
	}
	client, err := rcon.NewClient(conn, rconServerPassword)
	if err != nil {
		log.Fatal(err)
	}
	operator := rcon.NewOperator(client)
	get.NewUserListHandler(operator)
	post.NewUserItemHandler(operator)

	http.HandleFunc("/users", get.UserListHandler)
	http.HandleFunc("/users/", post.UserItemHandler)

	hostAddress := os.Getenv("HOST_ADDRESS")
	server := &http.Server{
		Addr:    hostAddress,
		Handler: nil,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			return
		}
	}()
	fmt.Println("call server")
	log.Fatal(server.ListenAndServe())
}

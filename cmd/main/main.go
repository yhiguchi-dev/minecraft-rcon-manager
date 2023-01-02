package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"rcon-client/internal/rcon"
	"strings"
	"time"
)

func main() {
	//command := flag.String("command", "", "command")
	//args := flag.String("args", "", "args")
	//dir := flag.String("dir", "", "dir")
	command := "java"
	args := "-jar ./spigot.jar nogui"
	dir := "/Users/yuki.higuchi/work/rcon-client/test"
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	cmd := exec.Command(command, strings.Fields(args)...)
	cmd.Dir = dir
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		scanner := bufio.NewScanner(io.MultiReader(stdoutPipe, stderrPipe))
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	go func() {
		<-ctx.Done()
		fmt.Println("interrupt!!")
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", "localhost", 25576), 3*time.Second)

		if err != nil {
			log.Fatal(err)
		}
		client, err := rcon.NewClient(conn, "asdf")
		if err != nil {
			log.Fatal(err)
		}
		err = client.Send("stop")
		if err != nil {
			log.Fatal(err)
		}
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	//host := flag.String("host", "", "host name")
	//port := flag.Int("port", 0, "port number")
	//pass := flag.String("pass", "", "rcon password for minecraft server")
	//exec := flag.String("exec", "", "exec command to minecraft server")
	//flag.Parse()
	//conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", *host, *port), 3*time.Second)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//client, err := rcon.NewClient(conn, *pass)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = client.Send(*exec)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

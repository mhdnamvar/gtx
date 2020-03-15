package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	port   int
	echo   bool
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "runs a gtx server",
		Long:  ``,
		Run:   checkRunFlags,
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&port, "port", "p", 8583, "Listening port")
	runCmd.Flags().BoolVarP(&echo, "echo-server", "e", false, "Run a echo server")
}

func checkRunFlags(cmd *cobra.Command, args []string) {
	if echo {
		echoServer(cmd, args)
	} else {
		_ = cmd.Usage()
	}

}

func echoServer(cmd *cobra.Command, args []string) {
	server, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if server == nil {
		log.Panicf("couldn't start listening: %v", err)
	} else {
		log.Print("gtx is listening on " + strconv.Itoa(port))
	}
	connections := acceptConnections(server)
	for {
		go handleConn(<-connections)
	}
}

func acceptConnections(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				log.Panicf("couldn't accept: %v", err)
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		_, err = client.Write(append([]byte("echo@gtx: "), line...))
		if err != nil {
			log.Fatal(err)
		}
	}
}

package cmd

import (
	"bufio"
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
		log.Fatalf("couldn't start listening: %v", err)
	}

	log.Printf("gtx is listening on %s", strconv.Itoa(port))
	for {
		client, err := server.Accept()
		if client == nil {
			log.Fatalf("couldn't accept: %v", err)
		}
		go func(c net.Conn) {
			log.Printf("%v connected", c.RemoteAddr())
			b := bufio.NewReader(c)
			for {
				line, err := b.ReadBytes('\n')
				if err != nil {
					log.Printf("%v disconnected", c.RemoteAddr())
					return
				}
				log.Print(string(line))
				_, err = c.Write(append([]byte("echo@gtx: "), line...))
				if err != nil {
					log.Print(err)
					return
				}
			}
		}(client)
	}
}
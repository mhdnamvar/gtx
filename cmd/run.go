package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var (
	port   int
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "runs a gtx server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&port, "port", "p", 8583, "Listening port")
}

func run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		log.Println("gtx server is listening to", port)
		_, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		break
	}
}

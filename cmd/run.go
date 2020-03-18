package cmd

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	echoPort int
	restPort int
	echo     bool
	rest     bool
	all      bool
	wg       sync.WaitGroup
	runCmd   = &cobra.Command{
		Use:   "run",
		Short: "runs a gtx server",
		Long:  ``,
		Run:   checkRunFlags,
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&echoPort, "echo-server-port", "", 8583, "Echo server listening port")
	runCmd.Flags().IntVarP(&restPort, "rest-server-port", "", 8584, "Rest api server listening port")
	runCmd.Flags().BoolVarP(&echo, "echo-server", "e", false, "Run a echo server")
	runCmd.Flags().BoolVarP(&rest, "rest-server", "r", false, "Run a rest api server")
	runCmd.Flags().BoolVarP(&all, "all", "a", false, "Run all servers")
}

func checkRunFlags(cmd *cobra.Command, args []string) {
	if echoPort == restPort {
		log.Fatal("The ports should be unique")
	}

	if all {
		wg.Add(1)
		go echoServer(cmd, args)
		wg.Add(1)
		go restServer(cmd, args)
	} else {
		if echo {
			wg.Add(1)
			go echoServer(cmd, args)
		}

		if rest {
			wg.Add(1)
			go restServer(cmd, args)
		}
	}

	if !all && !echo && !rest {
		_ = cmd.Usage()
	}

	wg.Wait()
}

func echoServer(cmd *cobra.Command, args []string) {
	defer wg.Done()
	server, err := net.Listen("tcp", ":"+strconv.Itoa(echoPort))
	if server == nil {
		log.Fatalf("couldn't start listening: %v", err)
	}

	log.Printf("gtx echo server is listening on %d", echoPort)
	for {
		client, err := server.Accept()
		if client == nil {
			log.Fatalf("couldn't accept: %v", err)
		}
		wg.Add(1)
		go func(c net.Conn) {
			defer wg.Done()
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

func restServer(cmd *cobra.Command, args []string) {
	defer wg.Done()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New() // gin.Default() for logging requests
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "gtx rest api"})
	})
	log.Printf("gtx rest api is listening on %d", restPort)
	r.Run(":" + strconv.Itoa(restPort))
}

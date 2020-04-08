package main

import (
	"bufio"
	"log"
	"log/syslog"
	"net"
	"os"
	"strings"

	"github.com/aTTiny73/multilogger/logs"
	_ "github.com/go-sql-driver/mysql"
)

// handleConnections heandles a tcp connection from client
func handleConnection(connection net.Conn, log *logs.MultipleLog) {

	log.Infof("%s just Connected. \n", connection.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			log.Error("Listening of:" + connection.RemoteAddr().String() + " stopped.")
			return
		}

		log.Info(connection.RemoteAddr().String() + " Says : " + strings.TrimSpace(string(netData)))

		connection.Write([]byte(string("Server: Message recived \n")))

		defer connection.Close()
	}

}

func main() {

	fileLog1 := logs.NewFileLogger("ServerLog")
	defer fileLog1.Close()

	syslog, _ := logs.NewSysLogger(syslog.LOG_NOTICE, log.LstdFlags)

	stdlog := logs.NewStdLogger()
	defer stdlog.Close()

	//databaseLog := logs.NewDataBaseLog(logs.DatabaseConfiguration())

	log := logs.NewCustomLogger(false, fileLog1, stdlog, syslog /*,databaseLog*/)

	arguments := os.Args
	if len(arguments) == 1 {
		log.Warn("Port not provieded...Unable to run.")
		return
	}

	PORT := ":" + arguments[1]
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Error(err)
		return
	}

	defer listener.Close()

	log.Info("Ready to listen...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err)
			return
		}

		//Starts a new goroutine each time it has to serve a TCP client.
		go handleConnection(conn, log)
	}
}

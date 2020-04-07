package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]

	conn, _ := net.Dial("tcp", PORT)
	defer conn.Close()

	//Client to server
	for {
		fmt.Println("Message to server:")
		inputscan := bufio.NewScanner(os.Stdin)
		inputscan.Scan()
		message := inputscan.Text()
		fmt.Fprintf(conn, message+"\n")

		//Server response to client
		serverRespons, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Server droped.")
			return
		}
		fmt.Println(serverRespons)
	}
}

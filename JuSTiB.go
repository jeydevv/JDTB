package main

// Imports all modules used
import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Welcome message
	wm := `
	+-+-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+
	|j|e|y|d|e|v|v|s| |T|o|o|l|b|o|x|
	+-+-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+
	`
	// Help message
	hm := `
Usage:
	encrypt    <MESSAGE>          <KEY>     (Encrypts a string using the Caesar Cipher)
	decrypt    <MESSAGE>  		  <KEY>     (Decrypts a string using the Caesar Cipher)
	checkport  <PROTOCOL>         <PORT>    (Checks if a port is open)
	customping <PROTOCOL:IP:PORT> <MESSAGE> (Send a raw text message to a specified IP & port)
	filedelete <DIRECTORY>        <FILE>    (Deletes a file from CWD and all subdirectories)
	welcome                       		    (Shows the welcome message)
	clear                         		    (Clears the console)
	help                          		    (Show this page, hello)
	exit                          		    (Exits the program)
	`

	fmt.Println(wm)

	// Infinate loop that gets input
	exit := false
	for exit == false {
		fmt.Print(">: ")

		// Variable declaration then scanning in input
		var command string
		var subcommand string
		var subcommand2 string
		fmt.Scanln(&command, &subcommand, &subcommand2)

		// If nothing is entered as the initial command it just continues the loop
		if command == "" {
			continue
		}

		// Set strings to lower for easier processing
		command = strings.ToLower(command)
		subcommand = strings.ToLower(subcommand)
		subcommand2 = strings.ToLower(subcommand2)

		// Main case statement to handles main command
		switch strings.ToLower(command) {
		case "clear":
			ClearCMD()
			break
		case "exit":
			ExitCMD()
			exit = true
			break
		case "checkport":
			fmt.Println(CheckPortCMD(subcommand, subcommand2))
			break
		case "encrypt":
			fmt.Println(EncryptCMD(subcommand, subcommand2))
			break
		case "decrypt":
			fmt.Println(DecryptCMD(subcommand, subcommand2))
			break
		case "help":
			fmt.Println(hm)
			break
		case "welcome":
			fmt.Println(wm)
			break
		case "filedelete":
			fmt.Println(FiledeleteCMD(subcommand, subcommand2))
			break
		case "customping":
			fmt.Println(CustomPingCMD(subcommand, subcommand2))
		default:
			fmt.Println("Unrecognised command, for help enter 'help'")
			break
		}
	}
}

// ClearCMD clears the console
func ClearCMD() {
	// Uses "cls" to clear console
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// ExitCMD prints exit message
func ExitCMD() {
	// Prints "Terminating~\n"
	fmt.Println("Terminating~")
	fmt.Println()
}

// CheckPortCMD checks a port
func CheckPortCMD(protocol, port string) string {
	// "Dials" the specified port on local PC and if it didn't throw an error it's open
	// Also uses specified protocol
	connection, err := net.Dial(protocol, ":"+port)

	if err != nil {
		return err.Error()
	}
	defer connection.Close()

	return "Port is open"
}

// EncryptCMD is the command used to encrypt a string using the Caesar Cipher, not quite working
func EncryptCMD(message, key string) string {

	return "doesnt work yet"

	intKey, err := strconv.Atoi(key)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	if err != nil {
		return err.Error()
	}

	var encryptedMessage string

	for i := 0; i < len(message); i++ {
		for t := 0; t < len(alphabet); t++ {
			if strings.ToLower(string(message[i])) == string(alphabet[t]) {
				if intKey > 25 || intKey < 1 {
					return "Key error"
				}

				if t-intKey >= 0 {
					encryptedMessage += string(alphabet[t-intKey])
				} else {
					encryptedMessage += string(alphabet[len(alphabet)-(intKey-t)])
				}
				break
			}
		}
	}
	return encryptedMessage
}

// DecryptCMD is the command used to decrypt a string using the Caesar Cipher, not quite working
func DecryptCMD(message, key string) string {

	return "doesnt work yet"

	intKey, err := strconv.Atoi(key)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	if err != nil {
		return err.Error()
	}

	var decryptedMessage string

	for i := 0; i < len(message); i++ {
		for t := 0; t < len(alphabet); t++ {
			if strings.ToLower(string(message[i])) == string(alphabet[t]) {
				if intKey > 25 || intKey < 1 {
					return "Key error"
				}

				if t-intKey >= 0 {
					decryptedMessage += string(alphabet[t+intKey])
				} else {
					decryptedMessage += string(intKey - (len(alphabet) + t))
				}

				break
			}
		}
	}
	return decryptedMessage
}

var count int // Needs to be public due to recursion

// FiledeleteCMD is used to delete files form a directory and all subdirectories
func FiledeleteCMD(directory, file string) string {
	// Gets a list of all files in CWD (current working directory)
	files, err := ioutil.ReadDir(directory)
	// If there is an error, return the error
	if err != nil {
		return err.Error()
	}

	// Uses recursion to acces subdirectories and deletes all files with the given filename
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			FiledeleteCMD(directory+"/"+files[i].Name(), file)
		}

		if files[i].Name() == file {
			err = os.Remove(string(directory + "/" + files[i].Name()))

			// If the file couldn't be removed notify user
			if err != nil {
				fmt.Println("Failed to delete a file")
			}
			count++ // Keeps track of how many files have been deleted
		}
	}
	return strconv.Itoa(count) + " File(s) successfully deleted"
}

// CustomPingCMD is used to send a rawtext massage to a specified IP through a specified port
func CustomPingCMD(fullIP, message string) string {
	var protocol string

	// Splits the protocol and the IP
	for i := 0; i < len(fullIP); i++ {
		if string(fullIP[i]) == ":" {
			protocol = fullIP[:i]
			fullIP = fullIP[i+1:]
			break
		}
	}

	// Attempts a connection
	connection, err := net.Dial(protocol, fullIP)

	// If an error is thown it is displayed to user
	if err != nil {
		return err.Error()
	}
	defer connection.Close()

	fmt.Fprintf(connection, message)

	// A success message is outputted
	return "Success: " + message + " sent to " + fullIP + " through " + protocol
}

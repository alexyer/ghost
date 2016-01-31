package cli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/alexyer/ghost/client"
)

var (
	helpMessage  = initHelpMessage()
	regularState *terminal.State
)

// initialize endless cli-session with provided client
// as a connection to ghost-server
func StartCliSession(c *client.GhostClient) {
	prepareCliSession()

	for {
		comm, args, err := processUserInput()
		if err != nil {
			log.Printf("Error on input processing: %s", err.Error())
			continue
		}

		result, err := makeRequest(c, comm, args)
		if err != nil {
			log.Printf("Error on request: %s", err.Error())
			continue
		}

		fmt.Println(result)
	}
}

func makeRequest(c *client.GhostClient, comm string, args []string) (string, error) {
	switch comm {
	case "PING":
		return pingServer(c, args)
	case "SET":
		if err := setValue(c, args); err != nil {
			return "", err
		}
	case "GET":
		return getValue(c, args)
	case "DEL":
		if err := delValue(c, args); err != nil {
			return "", err
		}
	case "CGET":
		if err := getColl(c, args); err != nil {
			return "", err
		}
	case "CADD":
		if err := addColl(c, args); err != nil {
			return "", err
		}
	case "HELP", "/H":
		return help(c, args)
	case "EXIT", "QUIT", "/Q":
		closeCliSession()
	default:
		return "", errors.New("unknown command: " + comm)
	}

	return "OK", nil
}

func prepareCliSession() {
	var err error

	regularState, err = terminal.MakeRaw(0)
	if err != nil {
		log.Fatalf("Error on cli session startup: %s. Exiting.", err.Error())
	}

	log.Println("Cli-ghost session was started. Type 'help' or '/h' for more information.")
}

// util cli-close handler
func closeCliSession() {
	terminal.Restore(0, regularState)

	log.Print("Cli session was successfully closed.")
	os.Exit(0)
}

// util help handler
func help(c *client.GhostClient, args []string) (string, error) {
	if len(args) != 0 {
		return "", errors.New(fmt.Sprintf("wrong number of arguments for HELP: need 0, get %d", len(args)))
	}

	return helpMessage, nil
}

func initHelpMessage() string {
	hm := []string{
		"Welcome to ghost-cli tool. To see this message enter '/h' or 'help'.",
		"To exit ghost-cli tool enter 'exit', 'quit' or '/q'.",
		"",
		"To add some value to collection enter 'set key value'",
		"To get some value from collection enter 'get key'",
		"To delete some value from collection enter 'del key'",
		"",
		"To a new collection to storage enter 'cadd key'",
		"To select some collection in storage enter 'cget key'",
	}

	return strings.Join(hm, "\n")
}

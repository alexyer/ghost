package cli

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

const CLI_GREETING = "> "

var term = terminal.NewTerminal(os.Stdin, CLI_GREETING)

func processUserInput() (string, []string, error) {
	commStr, err := readUserInput()
	if err != nil {
		return "", nil, errors.New("error on read input: " + err.Error())
	}

	return parseCommand(commStr)
}

func readUserInput() (string, error) {
	return term.ReadLine()
}

func parseCommand(commStr string) (string, []string, error) {
	comm, args, err := parseCommandString(commStr)
	if err != nil {
		return "", nil, errors.New("error on command parsing: " + err.Error())
	}

	arguments := make([]string, 0)
	for _, str := range args {
		arguments = append(arguments, strings.TrimSpace(str))
	}

	return strings.ToUpper(strings.TrimSpace(comm)), arguments, err
}

func parseCommandString(commStr string) (string, []string, error) {
	var args []string
	var err error

	warnQuoteEscaping(commStr)
	if strings.Contains(commStr, " \"") && !strings.Contains(commStr, "\\\"") {
		if args, err = splitQuotedCommand(commStr); err != nil {
			return "", nil, err
		}
	} else {
		args = strings.Split(commStr, " ")
	}

	return args[0], args[1:], nil
}

func splitQuotedCommand(commStr string) ([]string, error) {
	// Check if all quotes in the command string matches.
	// If the number of quote characters is odd - there are unmatched quotes.
	if (strings.Count(commStr, "\"") & 1) != 0 {
		return nil, errors.New("wrong syntax: unmatched quotes")
	}

	re := regexp.MustCompile(`\w+|"[\w ]*"`)

	args := re.FindAllString(commStr, -1)

	for i := range args {
		if args[i] == `""` {
			return nil, errors.New("empty argument")
		}
		args[i] = strings.Replace(args[i], "\"", "", -1)
	}

	return args, nil
}

// TODO (nikitasmall): make hard analysis to aloid this
func warnQuoteEscaping(commStr string) {
	if strings.Contains(commStr, "\\\"") {
		log.Println("Escaped quotes is not supported for now.")
	}
}

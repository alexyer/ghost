package cli

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func processUserInput() (string, []string, error) {
	commStr, err := readUserInput()
	if err != nil {
		return "", nil, errors.New("error on read input: " + err.Error())
	}

	return parseCommand(commStr)
}

func readUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	return reader.ReadString('\n')
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
		args, err = splitQuotedCommand(commStr)
		if err != nil {
			return "", nil, err
		}
	} else {
		args = strings.Split(commStr, " ")
		if len(args) > 3 {
			return "", nil, errors.New("too many arguments")
		}
	}

	return args[0], args[1:], nil
}

// TODO(nikitasmall): refactor (smells like a black magic but it's only strings with quotes)
func splitQuotedCommand(commStr string) ([]string, error) {
	startQuoteInd, endQuoteInd := strings.Index(commStr, " \""), strings.Index(commStr, "\" ")

	// closing quote at the end of command
	if endQuoteInd == -1 {
		partTwo := strings.Replace(commStr[startQuoteInd+2:len(commStr)-1], "\"", "", -1)

		// only second argument is in quotes
		if strings.Count(commStr[:startQuoteInd+2], " ") > 1 {
			return append(strings.Split(commStr[:startQuoteInd], " "), partTwo), nil
		} else { // there only one argument (and it is in quotes)
			if len(partTwo) == 0 {
				return nil, errors.New("empty first argument")
			}

			return []string{commStr[:strings.Index(commStr, " ")], partTwo}, nil
		}
	} else {
		// two aguments, both or first are in the quotes
		partTwo := commStr[startQuoteInd+2 : endQuoteInd]
		if len(partTwo) == 0 {
			return nil, errors.New("empty first argument")
		}

		commStr = strings.Replace(commStr, partTwo, "", 1)
		args := strings.Split(commStr, "\"\"")
		return []string{args[0], partTwo, strings.Replace(args[1], "\"", "", 2)}, nil
	}
}

// TODO (nikitasmall): make hard analysis to aloid this
func warnQuoteEscaping(commStr string) {
	if strings.Contains(commStr, "\\\"") {
		log.Println("Escaped quotes is not supported for now.")
	}
}

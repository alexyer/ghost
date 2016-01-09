package main

import (
	"bufio"
	"errors"
	"fmt"
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

func splitQuotedCommand(commStr string) ([]string, error) {
	startQuoteInd, endQuoteInd := strings.Index(commStr, " \""), strings.Index(commStr, "\" ")

	if endQuoteInd == -1 {
		partTwo := commStr[startQuoteInd+2 : len(commStr)-1]

		return []string{commStr[:strings.Index(commStr, " ")], partTwo}, nil
	} else {
		partTwo := commStr[startQuoteInd+2 : endQuoteInd]
		if len(partTwo) == 0 {
			return nil, errors.New("empty second argument")
		}

		commStr = strings.Replace(commStr, partTwo, "", 1)
		args := strings.Split(commStr, "\"\"")
		return []string{args[0], partTwo, strings.Replace(args[1], "\"", "", 2)}, nil
	}
}

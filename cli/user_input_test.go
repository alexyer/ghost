package cli

import "testing"

// this test sequense shows how command parsing works

func TestParsePing(t *testing.T) {
	commString := "ping"

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "PING" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 0 {
		t.Error("Error on args parsing: args: ", args)
	}
}

// generic two-argument command
func TestBaseParseSet(t *testing.T) {
	commString := "set hello bye"

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "SET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 2 {
		t.Error("Error on args parsing: args: ", args)
	}
}

// generic one-argument command
func TestBaseParseGet(t *testing.T) {
	commString := "get hello"

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "GET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 1 {
		t.Error("Error on args parsing: args: ", args)
	}
}

func TestQuotedParseSet(t *testing.T) {
	commString := `set "hello" "bye"`

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "SET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 2 {
		t.Error("Error on args parsing: args: ", args)
	}
}

func TestBigQuotedParseSet(t *testing.T) {
	commString := `set "hello i love you" "can you tell me your name"`

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "SET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 2 {
		t.Errorf("Error on args parsing: args: %v, len: %d", args, len(args))
	}

	if args[0] != "hello i love you" || args[1] != "can you tell me your name" {
		t.Error("Error on args parsing in quotes: args: ", args)
	}
}

func TestQuotedParseGet(t *testing.T) {
	commString := "get \"hello\""

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "GET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 1 {
		t.Error("Error on args parsing: args: ", args)
	}
}

func TestBigQuotedParseGet(t *testing.T) {
	commString := "get \"people are strange\""

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "GET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 1 {
		t.Error("Error on args parsing: args: ", args)
	}

	if args[0] != "people are strange" {
		t.Error("Error on args quotes parsing: args: ", args)
	}
}

func TestOneQuotedPairParseSet(t *testing.T) {
	commString := "set song \"hello i love you\""

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "SET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 2 {
		t.Error("Error on args parsing: args: ", args)
	}

	if args[1] != "hello i love you" || args[0] != "song" {
		t.Error("Error on args parsing in quotes: args: ", args)
	}
}

func TestOneQuotedReversedPairParseSet(t *testing.T) {
	commString := "set \"hello i love you\" song"

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	if comm != "SET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 2 {
		t.Error("Error on args parsing: args: ", args)
	}

	if args[0] != "hello i love you" || args[1] != "song" {
		t.Error("Error on args parsing in quotes: args: ", args)
	}
}

func TestFirstArgumentEmptyParseSetGet(t *testing.T) {
	commString := "set \"\" song"

	_, _, err := parseCommand(commString)
	if err == nil {
		t.Error("Parsing of command with empty first argument doesn't raise an error (for set).")
	}

	commString = "get \"\""

	_, _, err = parseCommand(commString)
	if err == nil {
		t.Error("Parsing of command with empty first argument doesn't raise an error (for get).")
	}
}

func TestParseSetWithEscapedQuoteWarning(t *testing.T) {
	commString := "set quote \\\""

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on correct input: ", err.Error())
	}

	if comm != "SET" {
		t.Error("Error on command parsing: wrong command: ", comm)
	}

	if len(args) != 2 {
		t.Error("Error on args parsing: args: ", args)
	}

	if args[0] != "quote" || args[1] != "\\\"" {
		t.Error("Error on args parsing in quotes: args: ", args)
	}
}

func TestParseUnmatchedQuotes(t *testing.T) {
	commString := "set quote \""

	_, _, err := parseCommand(commString)
	if err == nil {
		t.Error("Ignoring error with unmatched quote")
	}
}

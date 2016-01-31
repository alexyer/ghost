package cli

import (
	"testing"
)

// this test sequence shows part with connection to server

func TestPing(t *testing.T) {
	result, err := makeRequest(c, "PING", []string{})
	if err != nil {
		t.Error("Error on PING command: ", err.Error())
	}

	if result != "Pong!" {
		t.Error("Wrong PING response: ", result)
	}
}

func TestSetGet(t *testing.T) {
	result, err := makeRequest(c, "GET", []string{"world"})
	if err == nil {
		t.Error("Storage doesn't rise error on non-existing key.")
	}

	result, err = makeRequest(c, "SET", []string{"world", "earth"})
	if err != nil {
		t.Error("Error on value inserting: ", err.Error())
	}

	result, err = makeRequest(c, "GET", []string{"world"})
	if err != nil {
		t.Error("Error on value retrieving: ", err.Error())
	}

	if result != "earth" {
		t.Error("Wrong value returned: ", result)
	}
}

func TestSetGetDel(t *testing.T) {
	result, err := makeRequest(c, "SET", []string{"world", "earth"})
	if err != nil {
		t.Error("Error on value inserting: ", err.Error())
	}

	result, err = makeRequest(c, "GET", []string{"world"})
	if err != nil {
		t.Error("Error on value retrieving: ", err.Error())
	}

	if result != "earth" {
		t.Error("Wrong value returned: ", result)
	}

	result, err = makeRequest(c, "DEL", []string{"world"})
	if err != nil {
		t.Error("Error on value deleting: ", err.Error())
	}

	result, err = makeRequest(c, "GET", []string{"world"})
	if err == nil {
		t.Error("Storage doesn't rise error on non-existing key.")
	}
}

func TestSetGetCollection(t *testing.T) {
	result, err := makeRequest(c, "CADD", []string{"aliens"})
	if err != nil {
		t.Error("Error on collection creation: ", err.Error())
	}

	result, err = makeRequest(c, "CGET", []string{"aliens"})
	if err != nil {
		t.Error("Error on collection selection: ", err.Error())
	}

	result, err = makeRequest(c, "GET", []string{"world"})
	if err == nil {
		t.Error("Storage doesn't rise error on non-existing key.")
	}

	result, err = makeRequest(c, "SET", []string{"world", "mars"})
	if err != nil {
		t.Error("Error on value inserting: ", err.Error())
	}

	result, err = makeRequest(c, "GET", []string{"world"})
	if err != nil {
		t.Error("Error on value retrieving: ", err.Error())
	}

	if result != "mars" {
		t.Error("Wrong value returned: ", result)
	}

	result, err = makeRequest(c, "CGET", []string{"main"})
	if err != nil {
		t.Error("Error on collection selection: ", err.Error())
	}

	result, err = makeRequest(c, "GET", []string{"world"})
	if err == nil {
		t.Error("Storage doesn't rise error on non-existing key.")
	}
}

func TestWrongCommand(t *testing.T) {
	_, err := makeRequest(c, "GODMODE", []string{"on"})
	if err == nil {
		t.Error("Unknown command doesn't raise an error (or godmode is on).")
	}
}

func TestTooManyArgsNum(t *testing.T) {
	_, err := makeRequest(c, "GODMODE", []string{"health", "mana", "weapon", "xp"})
	if err == nil {
		t.Error("Too big args number doesn't raise an error (or godmode is on).")
	}
}

func TestCommandsWithWrongArgsNum(t *testing.T) {
	_, err := makeRequest(c, "PING", []string{"fast"})
	if err == nil {
		t.Error("Wrong args number doesn't raise an error for PING!.")
	}

	_, err = makeRequest(c, "SET", []string{"conquer", "world", "earth"})
	if err == nil {
		t.Error("Wrong args number doesn't raise an error for SET.")
	}

	_, err = makeRequest(c, "GET", []string{"give", "world"})
	if err == nil {
		t.Error("Wrong args number doesn't raise an error for GET.")
	}

	_, err = makeRequest(c, "DEL", []string{})
	if err == nil {
		t.Error("Wrong args number doesn't raise an error for DEL.")
	}

	_, err = makeRequest(c, "CADD", []string{"brave", "new"})
	if err == nil {
		t.Error("Wrong args number doesn't raise an error for CADD.")
	}

	_, err = makeRequest(c, "CGET", []string{"cruel", "world"})
	if err == nil {
		t.Error("Wrong args number doesn't raise an error for CGET.")
	}
}

func TestCollectionErrorState(t *testing.T) {
	_, err := makeRequest(c, "CADD", []string{"middlearth"})
	if err != nil {
		t.Error("Error on collection creation: ", err.Error())
	}

	_, err = makeRequest(c, "CADD", []string{"middlearth"})
	if err == nil {
		t.Error("Error doesn't raise on collection doubles")
	}

	_, err = makeRequest(c, "CGET", []string{"numenor"})
	if err == nil {
		t.Error("Error doesn't raise on select non-existed collection")
	}
}

func TestHelp(t *testing.T) {
	result, err := makeRequest(c, "HELP", []string{})
	if err != nil {
		t.Error("Error on HELP command: ", err.Error())
	}

	if len(result) < len(helpMessage) {
		t.Errorf("Wrong HELP response, message too small, result: %s", result)
	}
}

func TestHelpWrong(t *testing.T) {
	_, err := makeRequest(c, "HELP", []string{"cavabanga"})
	if err == nil {
		t.Error("Help ignoring wrong argument number.")
	}
}

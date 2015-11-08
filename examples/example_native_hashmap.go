package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	testHash := make(map[string]string)
	var words []string

	raw, err := ioutil.ReadFile("/usr/share/dict/cracklib-small")

	if err == nil {
		words = strings.Split(string(raw), "\n")
	}

	for _, word := range words {
		testHash[word] = word
	}
}

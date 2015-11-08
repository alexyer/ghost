package main

import (
	"io/ioutil"
	"strings"

	"github.com/alexyer/ghost"
)

func main() {
	testHash := ghost.NewHashMap()
	var words []string

	raw, err := ioutil.ReadFile("/usr/share/dict/cracklib-small")

	if err == nil {
		words = strings.Split(string(raw), "\n")
	}

	for _, word := range words {
		testHash.Set(word, word)
	}
}

package main

import (
	"testing"

	"github.com/alexyer/ghost"
)

var storage *ghost.Storage

func init() {
	storage = ghost.GetStorage()
}

func TestStorage(t *testing.T) {
	storage := ghost.GetStorage()

	if storage == nil {
		t.Errorf("Expected storage instance. Got nil")
	}

	mainCollection := storage.GetCollection("main")

	if mainCollection == nil {
		t.Errorf("Expected collection. Got nil")
	}

	storage.AddCollection("newCollection")

	newCollection := storage.GetCollection("newCollection")

	if newCollection == nil {
		t.Errorf("Expected newCollection collection. Got nil")
	}

	storage.DelCollection("newCollection")

	newCollection = storage.GetCollection("newCollection")

	if newCollection != nil {
		t.Errorf("Expected nil. Got collection")
	}

	storage.AddCollection("existentCollection")

	if _, err := storage.AddCollection("existentCollection"); err == nil {
		t.Errorf("Expected 'collection exists' error. Got nil")
	}
}

func TestCollection(t *testing.T) {
	testCollection, _ := storage.AddCollection("testCollection")

	testCollection.Set("newkey", "42")

	if r, _ := testCollection.Get("newkey"); r != "42" {
		t.Errorf("Expected value. Got error")
	}

	testCollection.Set("unknownkey", "73")
	testCollection.Del("unknownkey")

	if _, err := testCollection.Get("unknownkey"); err == nil {
		t.Errorf("Expected nil. Got value")
	}

}

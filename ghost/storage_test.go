package ghost

import "testing"

var testStorage *Storage

func init() {
	testStorage = GetStorage()
}

func TestStorage(t *testing.T) {
	newStorage := GetStorage()

	if newStorage == nil {
		t.Errorf("Storage construtor is wrong.")
	}

	if GetStorage() != GetStorage() {
		t.Errorf("Storage is not singleton")
	}
}

func TestStorageGet(t *testing.T) {
	mainCollection := testStorage.GetCollection("main")

	if mainCollection == nil {
		t.Errorf("Storage get is wrong.")
	}

	unknownCollection := testStorage.GetCollection("unknown")

	if unknownCollection != nil {
		t.Errorf("Get unknown collection. Expected nil.")
	}
}

func TestStorageAdd(t *testing.T) {
	testStorage.AddCollection("test")

	if testStorage.GetCollection("test") == nil {
		t.Errorf("Expected *collection, got nil.")
	}
}

func TestStorageDel(t *testing.T) {
	testStorage.AddCollection("test")
	testStorage.DelCollection("test")

	if testStorage.GetCollection("test") != nil {
		t.Errorf("Expected nil, got *collection.")
	}
}

// Storage is a heart of ghost.
// Storage contains all collections and interaction methods.
// Storage is singleton.
package ghost

import "errors"

type Storage struct {
	collections map[string]*Collection
}

var storageInstance *Storage = nil

// Return Storage instance.
func GetStorage() *Storage {
	if storageInstance == nil {
		storageInstance = &Storage{}
		storageInstance.collections = make(map[string]*Collection)
		storageInstance.collections["main"] = newCollection("main")
	}

	return storageInstance
}

// Get collection from the Storage.
// Return *colleciton or nil.
func (s *Storage) GetCollection(name string) *Collection {
	return s.collections[name]
}

// Add new collection to the Storage.
func (s *Storage) AddCollection(name string) (*Collection, error) {
	if s.collections[name] != nil {
		return nil, errors.New("Collection already exists")
	}

	s.collections[name] = newCollection(name)

	return s.collections[name], nil
}

// Delete collection from the Storage.
func (s *Storage) DelCollection(name string) {
	delete(s.collections, name)
}

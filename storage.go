// Storage is a heart of ghost.
// Storage contains all collections and interaction methods.
// Storage is singleton.
package ghost

import "errors"

type storage struct {
	collections map[string]*collection
}

var storageInstance *storage = nil

// Return storage instance.
func Storage() *storage {
	if storageInstance == nil {
		storageInstance = &storage{}
		storageInstance.collections = make(map[string]*collection)
		storageInstance.collections["main"] = newCollection("main")
	}

	return storageInstance
}

// Get collection from the storage.
// Return *colleciton or nil.
func (s *storage) GetCollection(name string) *collection {
	return s.collections[name]
}

// Add new collection to the storage.
func (s *storage) AddCollection(name string) (*collection, error) {
	if s.collections[name] != nil {
		return nil, errors.New("Collection already exists")
	}

	s.collections[name] = newCollection(name)

	return s.collections[name], nil
}

// Delete collection from the storage.
func (s *storage) DelCollection(name string) {
	delete(s.collections, name)
}

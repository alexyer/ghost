// Collection is just wrapper around hashmap.
// Storage could contain multiple collections.
package ghost

type Collection struct {
	Name    string
	hashMap *hashMap
}

func newCollection(name string) *Collection {
	return &Collection{name, NewHashMap()}
}

func (c *Collection) Set(key, val string) {
	c.hashMap.Set(key, val)
}

func (c *Collection) Get(key string) (string, error) {
	return c.hashMap.Get(key)
}

func (c *Collection) Del(key string) {
	c.hashMap.Del(key)
}

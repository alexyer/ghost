// Collection is just wrapper around hashmap.
// Storage could contain multiple collections.
package ghost

type collection struct {
	Name    string
	hashMap *hashMap
}

func newCollection(name string) *collection {
	return &collection{name, NewHashMap()}
}

func (c *collection) Set(key, val string) {
	c.hashMap.Set(key, val)
}

func (c *collection) Get(key string) (string, error) {
	return c.hashMap.Get(key)
}

func (c *collection) Del(key string) {
	c.hashMap.Del(key)
}

package ghost

import "testing"

var c *Collection

func init() {
	c = newCollection("Test Collection")
	c.Set("Predefined Value", "73")
}

func TestCollectionNew(t *testing.T) {
	testName := "Test Collection"
	collection := newCollection(testName)

	if collection.Name != testName {
		t.Error("Wrong collection name. Expected: %s, Got: %s", testName, collection.Name)
	}
}

func TestCollectionSet(t *testing.T) {
	c.Set("Test Value", "42")

	if val, _ := c.Get("Test Value"); val != "42" {
		t.Error("Wrong Set. Expected: %s, Got: %s", "42", val)
	}
}

func TestCollectionGet(t *testing.T) {
	if val, _ := c.Get("Predefined Value"); val != "73" {
		t.Error("Wrong Get. Expected: %s, Got: %s", "73", val)
	}
}

func TestCollectionDel(t *testing.T) {
	c.Set("Test Value", "42")
	c.Del("Test Value")

	if _, err := c.Get("Test Value"); err == nil {
		t.Error("Wrong Del.")
	}
}

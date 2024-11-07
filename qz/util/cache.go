package util

// https://go.dev/play/p/j9elQATrszu

type Cache[K comparable, V any] struct {
	maxlen int
	keys   map[K]int
	vals   []V
}

func NewCache[k comparable, v any](len int) *Cache[k, v] {

	c := &Cache[k, v]{maxlen: len, keys: make(map[k]int)}
	return c
}

func (c *Cache[K, V]) Add(key K, val V) {
	if len(c.vals) >= c.maxlen {
		c.vals = c.vals[1:]
	}
	c.vals = append(c.vals, val)
	c.keys[key] = len(c.vals) - 1
}

func (c *Cache[K, V]) Get(key K) *V {
	if ndx, ok := c.keys[key]; ok {
		return &c.vals[ndx]
	}
	return nil
}

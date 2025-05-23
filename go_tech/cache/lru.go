package main

type DB interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
	Delete(key string) bool
}

type Cache struct {
	// Your implementation here
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return nil, false
}
func (c *Cache) Put(key string, value interface{}) {
	return
}
func (c *Cache) Delete(key string) bool {
	return false
}

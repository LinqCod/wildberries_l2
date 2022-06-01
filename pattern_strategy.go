package main

import "fmt"

type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct{}

func (al *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

type lru struct{}

func (al *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

type lfu struct{}

func (al *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) string {
	value := c.storage[key]
	delete(c.storage, key)
	return value
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("1", "a")
	cache.add("2", "b")

	lru := &lru{}
	cache.setEvictionAlgo(lru)

	_ = cache.get("2")
	cache.add("3", "c")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("4", "d")
	cache.add("5", "e")
}

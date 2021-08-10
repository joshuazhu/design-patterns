package main

func main() {
	f := &fifo{}

	cache := initCache(f)
	cache.add("1", "1")
	cache.add("2", "2")
	cache.add("3", "3")

	lfu := &lfu{}
	cache.setEvictionAlgo(lfu)
	cache.evict()

	lru := &lru{}

	cache.setEvictionAlgo(lru)
	cache.add("4", "4")
	cache.add("5", "5")
}

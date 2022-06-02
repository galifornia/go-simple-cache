package main

import (
	"github.com/galifornia/go-simple-cache/structures"
)

func main() {
	cache := structures.CreateCache()

	for _, val := range []string{"zebra", "pig", "cocrodile", "pig", "zebra", "orca", "leopard", "bison", "cocrodile", "gazelle", "lion", "orca", "tiger", "cow", "leopard", "pig", "purpoise"} {
		cache.Insert(val)
	}

	// Display ca
	cache.Queue.Display()
}

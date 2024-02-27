package main

import (
	"fmt"
	"sync"
	. "container/ring"
)

// Node class
type Node struct {
	id   string
	data map[string]string
	network *DHT
}

type DHT struct {
	hashring *Ring
	lock sync.RWMutex
}

func main() {
	fmt.Println("Hello World!")
	
}

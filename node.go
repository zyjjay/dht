package main

import (
	"container/ring"
	"crypto/sha256"
	"encoding/hex"
)

func (dht *DHT) init() {
	dht.hashring = ring.New(0)
}

func (dht *DHT) AddNode(node *Node) {
	dht.lock.Lock()
	defer dht.lock.Unlock()
	// Compute new hashring, then either append to current ring, or create a new one and populate (less efficient)
	// dht.hashring = append(dht.hashring, node)
	node.network = dht
}

func (dht *DHT) RemoveNode(id string) {
	dht.lock.Lock()
	defer dht.lock.Unlock()
}

func hash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

func (node *Node) Store(key string, val string) {
	hashedKey := hash(key)
	node.data[hashedKey] = val
}

func (node *Node) Get(key string) (string, bool) {
	hashedKey := hash(key)
	val, ok := node.data[hashedKey]
	return val, ok
}

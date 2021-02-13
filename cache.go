package main 

import (

 "fmt"
)

type Node struct {
	key string
	value string
	next *Node
	//prev *Node
}

type Cache struct {
	M map[string]*Node
	ListHead *Node
	NoOfKeys int
	Capacity int
	EvictionPolicy EvictionPolicyTypes
	//
}

func (c *Cache) InitCache(capacity int, evictionPolicy EvictionPolicyTypes) {
	//Init map - check
	(*c).M = make(map[string]*Node)
	(*c).Capacity = capacity
	(*c).EvictionPolicy = evictionPolicy
}

func (c *Cache) Get (key string) (string, error) {
	if c.CheckIfCacheEmpty() {
		return "", fmt.Errorf("Cache is empty")
	}

	//Check for node in map
	node, ok := (*c).M[key]
	if !ok {
		return "", fmt.Errorf("Key not present in map")
	}

	//Get value from list
	value := node.value

	//Move node to top of list
	c.moveNodeToTop(node)

	return value, nil

 }

func (c *Cache) Set (key string, value string) {
	//Check if capacity reached
	if (*c).NoOfKeys == (*c).Capacity {
		getEvictionPolicyImpl((*c).EvictionPolicy).RemoveElements(c)
	}

	//Check if value present in map
	_, ok := (*c).M[key]
	if !ok {
		c.addInCache(key, value)
	} else {
		c.updateInCache(key, value)
	}


 }

 func (c *Cache) Flush (key string) {

 }

 func (c *Cache) Delete(key string){

 }

 func (c *Cache) addInCache(key string, value string) {
	node := createNode(key, value)

	(*c).M[key] = node
	numberOfKeys := (*c).NoOfKeys
	(*c).NoOfKeys = numberOfKeys + 1
	
	head := (*c).ListHead
	if head == nil {
		head = node
		(*c).ListHead = head
	} else {
		for iter := (*c).ListHead; iter != nil; iter = iter.next {
			iter.next = node
			node.next = nil
			
		}
	}

	
	
 }

 func (c *Cache) updateInCache(key string, value string) {
	 node := (*c).M[key]

	 node.value = value

	 c.moveNodeToTop(node)
 }

 func (c *Cache) CheckIfCacheEmpty() bool {
	if (*c).NoOfKeys == 0 {
		return true
	}

	return false
 }

 func (c *Cache) moveNodeToTop(node *Node) {
	for iter := (*c).ListHead; iter.next != nil; iter = iter.next {
		if iter.next == node {
			iter.next = iter.next.next
		}
	}

	head := (*c).ListHead
	head.next = node
	node.next = (*c).ListHead
 }
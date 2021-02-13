package main

type EvictionPolicyImpl interface {
	RemoveElements(c *Cache)
}

type LRUImpl struct {}

func getEvictionPolicyImpl(evictionPolicy EvictionPolicyTypes) EvictionPolicyImpl {
	switch evictionPolicy {
	case LRU:
		return LRUImpl{}
	}

	return nil
}

func (impl LRUImpl)RemoveElements(c *Cache) {
	
	var prev *Node
	head := (*c).ListHead

	no:= (*c).NoOfKeys
	(*c).NoOfKeys = no - 1
	
	if head.next == nil {
		(*c).ListHead = nil
		delete((*c).M, head.key)
		return
	}

	for iter := (*c).ListHead; iter != nil; iter = iter.next {
		if iter.next == nil {
			prev.next = nil
			k := iter.key
			delete((*c).M, k)
			break
		}

		prev = iter
	}



}
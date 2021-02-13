package main




func createNode(key string, value string) *Node {
	node := &Node {
		key: key,
		value: value,
		next: nil,
	}

	return node
}
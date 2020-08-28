package datastructures

func buildLinkedList(nodeValues []interface{}) *LinkedList {
	l := &LinkedList{
		Size: len(nodeValues),
	}

	var lastNode *LLNode

	if len(nodeValues) > 0 {
		lastNode = &LLNode{Data: nodeValues[0]}
		l.Head = lastNode
		l.Tail = lastNode
	}

	for i := 1; i < len(nodeValues); i++ {
		newNode := &LLNode{Data: nodeValues[i]}
		lastNode.Next = newNode
		l.Tail = newNode
		lastNode = newNode
	}

	return l
}

func setupNodes(initData []interface{}) []*DLLNode {
	var initNodes []*DLLNode

	if len(initData) > 0 {
		for i := 0; i < len(initData); i++ {
			newNode := &DLLNode{Data: initData[i]}

			if i > 0 {
				prevNode := initNodes[i-1]
				prevNode.Next = newNode
				newNode.Prev = prevNode
			}

			initNodes = append(initNodes, newNode)
		}
	}

	return initNodes
}

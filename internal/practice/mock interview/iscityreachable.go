// do event in diff city from current location even if not there
// as long as it's "reachable"

type Player struct {
	Location *Node
}

type Node struct {
	Edges []*Node
}

func IsCityReachable(player *Player, destination *Node) (bool, error) {
	start := player.Location

	if start == nil {
		return false, NEWERROR
	}

	checkedNodes := map[*Node]bool

	currentNodes := []*Node{start}
	nextNodes := []*Node

	// load nodes into step array
	// check for dest, place into map
	// get next step
	// -

	getNextNodes := func(node *Node) []*Node {
		nodes := []*Node
		for _, n := range node.Edges {
			if checkedNodes[n] == false {
				nodes = append(nodes, n)
			}
		}

		return nodes
	}

	for len(currentNodes) > 0 {
		for _, node := range currentNodes {
			if node == destination {
				return true, nil
			}

			checkedNodes[node] == true
			nextNodes = append(nextNodes, getNextNodes(node))
		}

		currentNodes = nextNodes
		nextNodes = []*Node
	}

}

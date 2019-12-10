package exercises

import (
	"bufio"
	"io"
	"strings"
)

func Day6(r io.Reader) (int, int) {
	var c int
	g := NewGraph(r)
	Distance(g["COM"], 0, &c)
	d := Between(g["YOU"], g["SAN"])
	return c, d
}

type Node struct {
	Parent   *Node
	Children []*Node
}

func NewGraph(r io.Reader) map[string]*Node {
	scanner := bufio.NewScanner(r)
	graph := make(map[string]*Node)
	graph["COM"] = new(Node)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ")")
		parent, child := pair[0], pair[1]
		if _, ok := graph[parent]; !ok {
			graph[parent] = new(Node)
		}
		if _, ok := graph[child]; !ok {
			graph[child] = new(Node)
		}
		graph[child].Parent = graph[parent]
		graph[parent].Children = append(graph[parent].Children, graph[child])
	}
	return graph
}

func Distance(node *Node, d int, c *int) {
	*c = *c + d
	for _, n := range node.Children {
		Distance(n, d+1, c)
	}
}

func Between(node1 *Node, node2 *Node) int {
	var n1, n2 []*Node
	for n := node1.Parent; n != nil; n = n.Parent {
		n1 = append(n1, n)
	}
	for n := node2.Parent; n != nil; n = n.Parent {
		n2 = append(n2, n)
	}
	for {
		if n1[len(n1)-1] == n2[len(n2)-1] {
			n1, n2 = n1[:len(n1)-1], n2[:len(n2)-1]
		} else {
			return len(n1) + len(n2)
		}
	}
}

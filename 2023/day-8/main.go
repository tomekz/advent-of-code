package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	id    string
	left  *Node
	right *Node
}

type Network struct {
	nodes []*Node
}

func (n *Network) lookup(id string) *Node {
	for _, node := range n.nodes {
		if node.id == id {
			return node
		}
	}

	return nil
}

func main() {
	content, _ := os.ReadFile("input")
	lines := strings.Split(string(content), "\n")

	instructions := lines[0]

	nodes := lines[2:752]

	network := &Network{}

	for _, line := range nodes {
		matcher := regexp.MustCompile(`(\w{3})`)
		ids := matcher.FindAllString(line, -1)

		node := &Node{
			id: ids[0],
			left: &Node{
				id: ids[1],
			},
			right: &Node{
				id: ids[2],
			},
		}

		network.nodes = append(network.nodes, node)
	}

	currNode := network.lookup("AAA")
	found := false
	iterations := 0
	for !found {
		for i := 1; i <= len(instructions); i++ {
			instruction := instructions[i]

			if instruction == 'R' {
				currNode = network.lookup(currNode.right.id)
			} else {
				currNode = network.lookup(currNode.left.id)
			}

			fmt.Println(i, currNode.id)
			if currNode.id == "ZZZ" {
				fmt.Println("Found ZZZ ater ", i*iterations, " steps")
				found = true
				break
			}

			if i == len(instructions)-1 {
				i = 1
				iterations++
			}
		}
	}
}

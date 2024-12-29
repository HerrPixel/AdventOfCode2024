package day23

import (
	"slices"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() map[string][]string {
	nodes := make(map[string][]string)

	for _, l := range Tools.ReadByLines("./Day23/input.txt") {
		input := strings.Split(l, "-")

		_, exists := nodes[input[0]]
		if !exists {
			nodes[input[0]] = []string{}
		}

		nodes[input[0]] = append(nodes[input[0]], input[1])

		_, exists = nodes[input[1]]
		if !exists {
			nodes[input[1]] = []string{}
		}

		nodes[input[1]] = append(nodes[input[1]], input[0])
	}
	return nodes
}

// starting with any node that starts with "t", we consider pairs of nodes in its neighborhood and test for triangles
// We additionally store each triangle in a set, should we encounter doubles, i.e. two nodes starting with "t" in the same triangle
func Triangles() string {
	nodes := parseInput()

	found := make(map[string]bool, 0)

	for key, value := range nodes {
		if key[0] != 't' {
			continue
		}

		// tests each pair of nodes in the neighborhood of key
		for _, n := range value {
			for _, m := range value {

				// tests if n and m share an edge
				if slices.Contains(nodes[n], m) {

					// finally, find the unique sorted representation of this triangle and store it in a set
					tripple := []string{key, n, m}
					slices.Sort(tripple)
					hash := strings.Join(tripple[:], ",")

					_, exists := found[hash]
					if exists {
						continue
					}

					found[hash] = true
				}
			}
		}
	}

	return strconv.Itoa(len(found))
}

// Find the sorted, comma-seperated list of nodes in the largest clique in the graph
// We use the Bron-Kerbosch algorithm with pivoting, see below
func LargestClique() string {
	nodes := parseInput()
	keys := make([]string, 0, len(nodes))
	for k := range nodes {
		keys = append(keys, k)
	}

	maximalClique := BronKerbosch([]string{}, keys, []string{}, nodes)
	slices.Sort(maximalClique)

	return strings.Join(maximalClique, ",")
}

// Bron Kerbosch algorithm with pivoting to find all maximal clique of a graph
// Instead of saving each clique, we just keep track of the largest one, since that is what we are interested in
func BronKerbosch(R []string, P []string, X []string, nodes map[string][]string) []string {
	if len(P) == 0 && len(X) == 0 {
		return R
	}

	maximalClique := make([]string, 0)
	pivot := ""
	if len(P) == 0 {
		pivot = X[0]
	} else {
		pivot = P[0]
	}

	pivotNeighborhood := nodes[pivot]

	for _, v := range without(P, pivotNeighborhood) {
		vNeighborhood := nodes[v]

		res := BronKerbosch(
			union(R, v),
			intersection(P, vNeighborhood),
			intersection(X, vNeighborhood),
			nodes,
		)

		if len(res) > len(maximalClique) {
			maximalClique = res
		}

		P = without(P, []string{v})
		X = append(X, v)
	}

	return maximalClique
}

// helper functions to make bron-kerbosch more readable

// returns l ∪ {s}
func union(l []string, s string) []string {
	return append(Tools.Clone(l), s)
}

// returns l ∩ m
func intersection(l []string, m []string) []string {
	res := make([]string, 0)

	for _, s := range l {
		if slices.Contains(m, s) {
			res = append(res, s)
		}
	}

	return res
}

// returns l \ m
func without(l []string, m []string) []string {
	res := make([]string, 0)

	for _, s := range l {
		if !slices.Contains(m, s) {
			res = append(res, s)
		}
	}

	return res
}

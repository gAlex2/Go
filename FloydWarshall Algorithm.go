package main

import "math"

type Arch struct {
	start, end int
}

func floydWarshall(graph map[int][]int, weights map[Arch]int) [][]int {
	
	dist := make([][]int, len(graph))
	for i := 0; i < len(graph); i++ {
		dist[i] = make([]int, len(graph))
	}

	
	for i := range graph {
		for j := range graph {
			if i == j {
				dist[i][j] = 0
			} else if pointIsReachable(graph, i, j) {
				dist[i][j] = weights[Arch{i, j}]
			} else {
				dist[i][j] = math.MaxInt
			}
		}
	}

	
	for k := 1; k < len(graph); k++ {
		for i := 1; i < len(graph); i++ {
			for j := 0; j < len(graph); j++ {
				
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func pointIsReachable(graph map[int][]int, start, end int) bool {
	for _, v := range graph[start] {
		if v == end {
			return true
		}
	}
	return false
}

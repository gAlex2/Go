package main

import "math"

type Point struct {
	x, y int
}

func dijkstra(grid [][]int, cur Point) map[Point]int {

	distances := make(map[Point]int)
	queue := make(map[Point]int)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			distances[Point{x, y}] = math.MaxInt
			queue[Point{x, y}] = math.MaxInt

			
			if x == cur.x && y == cur.y {
				distances[Point{x, y}] = 0
				queue[Point{x, y}] = 0
			}
		}
	}

	for len(queue) > 0 {

		var min int = math.MaxInt
		var cur Point
		for k := range queue {
			if distances[k] < min {
				min = distances[k]
				cur = k
			}
		}

		delete(queue, cur)

		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				
				if x == 0 && y == 0 {
					continue
				}

				v := Point{cur.x + x, cur.y + y}

				weight := 1

				if (distances[cur] + weight) < distances[v] {
					distances[v] = distances[cur] + weight
				}
			}
		}

		unreachable := true
		for cLeft := range queue {
			if distances[cLeft] != math.MaxInt {
				unreachable = false
			}
		}
		if unreachable {
			break
		}
	}
	return distances
}

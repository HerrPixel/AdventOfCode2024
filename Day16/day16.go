package day16

import (
	"container/heap"
	"math"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type triple struct {
	x         int
	y         int
	direction int
}

type tuple struct {
	x int
	y int
}

func parseInput() ([][]bool, int, int, int, int) {
	grid := make([][]bool, 0)
	startX := 0
	startY := 0
	endX := 0
	endY := 0

	for i, s := range Tools.ReadByLines("./Day16/input.txt") {
		grid = append(grid, make([]bool, 0))

		for j, c := range strings.Split(s, "") {
			if c == "E" {
				endX = i
				endY = j
			}

			if c == "S" {
				startX = i
				startY = j
			}

			if c == "#" {
				grid[i] = append(grid[i], false)
			} else {
				grid[i] = append(grid[i], true)
			}
		}
	}

	return grid, startX, startY, endX, endY
}

func Part1() string {
	grid, startX, startY, endX, endY := parseInput()

	distances, _ := dijkstra(grid, startX, startY)

	minDistance := min(
		distances[endX][endY][0],
		distances[endX][endY][1],
		distances[endX][endY][2],
		distances[endX][endY][3],
	)

	return strconv.Itoa(minDistance)
}

func Part2() string {
	grid, startX, startY, endX, endY := parseInput()

	distances, predecessor := dijkstra(grid, startX, startY)

	BestPath := make(map[tuple]bool, 0)
	BestPath[tuple{endX, endY}] = true

	var q Tools.Queue[triple]

	minDistance := min(
		distances[endX][endY][0],
		distances[endX][endY][1],
		distances[endX][endY][2],
		distances[endX][endY][3],
	)

	for i := range 4 {
		if distances[endX][endY][i] == minDistance {
			q = Tools.Enqueue(q, triple{endX, endY, i})
		}
	}

	for !Tools.IsEmpty(q) {
		var f triple
		f, q, _ = Tools.Dequeue(q)

		BestPath[tuple{f.x, f.y}] = true

		for _, x := range predecessor[f.x][f.y][f.direction] {
			q = Tools.Enqueue(q, x)
		}
	}

	return strconv.Itoa(len(BestPath))
}

func dijkstra(grid [][]bool, startX int, startY int) ([][][4]int, [][][4][]triple) {
	pq := make(Tools.PriorityQueue[triple], 0)
	heap.Init(&pq)

	heap.Push(&pq, &Tools.Item[triple]{Value: triple{startX, startY, 1}, Priority: 0, Index: 0})

	distance := make([][][4]int, 0)
	predecessor := make([][][4][]triple, 0)

	moves := [4]tuple{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i, row := range grid {
		distance = append(distance, make([][4]int, 0))
		predecessor = append(predecessor, make([][4][]triple, 0))
		for range row {
			distance[i] = append(distance[i], [4]int{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64})
			predecessor[i] = append(predecessor[i], [4][]triple{{}, {}, {}, {}})
		}
	}

	distance[startX][startY][1] = 0

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*Tools.Item[triple])

		currDistance := curr.Priority
		x := curr.Value.x
		y := curr.Value.y
		direction := curr.Value.direction

		nextX := x + moves[direction].x
		nextY := y + moves[direction].y
		left := mod(direction-1, 4)
		right := mod(direction+1, 4)

		if distance[nextX][nextY][direction] == currDistance+1 {
			predecessor[nextX][nextY][direction] = append(predecessor[nextX][nextY][direction], triple{x, y, direction})
		}

		if distance[x][y][left] == currDistance+1000 {
			predecessor[x][y][left] = append(predecessor[x][y][left], triple{x, y, direction})
		}

		if distance[x][y][right] == currDistance+1000 {
			predecessor[x][y][right] = append(predecessor[x][y][left], triple{x, y, direction})
		}

		if distance[nextX][nextY][direction] > currDistance+1 && grid[nextX][nextY] {

			distance[nextX][nextY][direction] = currDistance + 1
			predecessor[nextX][nextY][direction] = []triple{{x, y, direction}}

			item := &Tools.Item[triple]{
				Value:    triple{nextX, nextY, direction},
				Priority: currDistance + 1,
				Index:    0,
			}

			heap.Push(&pq, item)
		}

		if distance[x][y][left] > currDistance+1000 {

			distance[x][y][left] = currDistance + 1000
			predecessor[x][y][left] = []triple{{x, y, direction}}

			item := &Tools.Item[triple]{
				Value:    triple{x, y, left},
				Priority: currDistance + 1000,
				Index:    0,
			}

			heap.Push(&pq, item)
		}

		if distance[x][y][right] > currDistance+1000 {

			distance[x][y][right] = currDistance + 1000
			predecessor[x][y][right] = []triple{{x, y, direction}}

			item := &Tools.Item[triple]{
				Value:    triple{x, y, right},
				Priority: currDistance + 1000,
				Index:    0,
			}

			heap.Push(&pq, item)
		}

	}

	return distance, predecessor
}

func mod(n int, m int) int {
	return (n%m + m) % m
}

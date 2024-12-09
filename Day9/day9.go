package day9

import (
	"strconv"
	"strings"

	"git.jonasseiler.de/Jonas/AdventOfCode2024/Tools"
)

func parseInput() []int {
	sizes := make([]int, 0)
	length := 0

	for _, c := range strings.Split(Tools.Read("./Day9/input.txt"), "") {
		n, _ := strconv.Atoi(c)
		sizes = append(sizes, n)
		length += n
	}

	disk := make([]int, length)
	id := 0
	isFreeSpace := false
	currPos := 0

	for _, n := range sizes {
		entry := -1
		if !isFreeSpace {
			entry = id
		}

		for i := range n {
			disk[currPos+i] = entry
		}

		currPos += n

		if !isFreeSpace {
			id++
		}

		isFreeSpace = !isFreeSpace
	}

	return disk
}

func Part1() string {
	disk := parseInput()
	index := len(disk) - 1
	i := 0

	for isInBounds(disk, index) && index >= i {
		if !isFreeSpace(disk, index) {
			disk, i = moveFileWithStartValue(disk, index, i)
		}
		index--
	}

	return strconv.Itoa(calculateChecksum(disk))
}

func Part2() string {
	disk := parseInput()
	index := len(disk) - 1

	for isInBounds(disk, index) {
		length := reverseLengthOfStrip(disk, index)

		if !isFreeSpace(disk, index) {
			disk, _ = moveFile(disk, index-(length-1))
		}
		index -= length
	}

	return strconv.Itoa(calculateChecksum(disk))
}

func calculateChecksum(l []int) int {
	total := 0
	for i, v := range l {
		if v == -1 {
			continue
		}

		total += i * v
	}

	return total
}

func moveFile(l []int, index int) ([]int, int) {
	return moveFileWithStartValue(l, index, 0)
}

func moveFileWithStartValue(l []int, index int, start int) ([]int, int) {
	length := lengthOfStrip(l, index)
	id := l[index]

	i := findFreeSpace(l, length, start)

	if i >= index || i == -1 {
		return l, start
	}

	for j := range length {
		l[i+j] = id
		l[index+j] = -1
	}

	return l, i + length
}

func isFreeSpace(l []int, index int) bool {
	return l[index] == -1
}

func isInBounds(l []int, index int) bool {
	return 0 <= index && index < len(l)
}

func reverseLengthOfStrip(l []int, index int) int {
	id := l[index]
	length := 1

	for isInBounds(l, index-length) && l[index-length] == id {
		length++
	}

	return length
}

func lengthOfStrip(l []int, index int) int {
	id := l[index]
	length := 1

	for isInBounds(l, index+length) && l[index+length] == id {
		length++
	}

	return length
}

func findFreeSpace(l []int, size int, start int) int {
	i := start
	for isInBounds(l, i) {
		length := lengthOfStrip(l, i)

		if l[i] == -1 && length >= size {
			return i
		}

		i += length
	}

	return -1
}

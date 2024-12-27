package day9

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
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

// We move files with the standard moving algorithm from part 2 but pretending all files have size 1
func Part1() string {
	disk := parseInput()
	index := len(disk) - 1
	i := 0

	for isInBounds(disk, index) && index >= i {
		if !isFreeSpace(disk, index) {
			// the pretending actually happens here where we say that the file starts at the same index it ends
			disk, i = moveFileWithStartValue(disk, index, i)
		}
		index--
	}

	return strconv.Itoa(calculateChecksum(disk))
}

// we move files with the described moving algorithm from the end of the list to the beginning
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

// moves the file starting at index to the next continuous free space starting from the beginning of the disk
func moveFile(l []int, index int) ([]int, int) {
	return moveFileWithStartValue(l, index, 0)
}

// moves the file starting at index to the next continous free space starting at start
// is the same as moveFile but we supply a start value to search from, to speed up the search when the beginning is mostly occupied
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

// Counts the length of the countinous strip ending at index
func reverseLengthOfStrip(l []int, index int) int {
	id := l[index]
	length := 1

	for isInBounds(l, index-length) && l[index-length] == id {
		length++
	}

	return length
}

// Counts the length of the continous space starting at index
func lengthOfStrip(l []int, index int) int {
	id := l[index]
	length := 1

	for isInBounds(l, index+length) && l[index+length] == id {
		length++
	}

	return length
}

// finds the next continuous free space of length at least size and index starting at start
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

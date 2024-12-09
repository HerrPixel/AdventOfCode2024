package day9

import (
	"fmt"
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
		if isFreeSpace {
			for i := range n {
				disk[currPos+i] = -1
			}
			currPos += n
		} else {
			for i := range n {
				disk[currPos+i] = id
			}
			currPos += n
			id++
		}

		isFreeSpace = !isFreeSpace
	}

	//fmt.Println(disk)

	return disk
}

func Part1() string {
	disk := parseInput()

	left := 0
	right := len(disk) - 1

	for left < right && right > 0 {
		if disk[right] != -1 {
			for disk[left] != -1 && left < len(disk)-1 {
				left++
			}

			if left < right {
				disk[left] = disk[right]
				disk[right] = -1
			}
		}

		right--
	}

	return strconv.Itoa(calculateChecksum(disk))
}

func Part2() string {
	disk := parseInput()

	fmt.Println(disk)

	index := len(disk) - 1

	for isInBounds(disk, index) {
		size := 1

		if !isFreeSpace(disk, index) {
			//fmt.Println("test")
			id := disk[index]

			size, _ = findLengthOfStrip(disk, id, index)
			//fmt.Println("id", id, "size", size)

			disk = moveFile(disk, id, size, index)
			//fmt.Println("disk:", disk)
			//disk[right] = -1

			/*
				for right-blockSize >= 0 {
					if disk[right-blockSize] == id {
						//disk[right-blockSize] = -1
						blockSize++
					} else {
						break
					}
				}

				fmt.Println("disk:", disk)
				fmt.Println("moving block of", id, "with size", blockSize)
				disk = moveFile(disk, id, blockSize, right-blockSize)
			*/
		}
		index -= size
	}

	fmt.Println(disk)

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

func moveFile(l []int, id int, size int, index int) []int {

	i := findFreeSpace(l, size)

	//fmt.Println("free space a")

	if i >= index {
		return l
	}

	//fmt.Println("test3")

	if i == -1 {
		return l
	}

	for j := range size {
		l[i+j] = id
		l[index-j] = -1
	}

	return l

	/*

		for index := range l {
			if l[index] == -1 {
				freeSpaceSize := 1

				for freeSpaceSize < size {
					if index+freeSpaceSize < len(l) {
						if l[index+freeSpaceSize] == -1 {
							freeSpaceSize++
						} else {
							break
						}
					} else {
						break
					}
				}

				if freeSpaceSize < size {
					continue
				}

				for i := range size {
					l[index+i] = id
				}

				break
			}
		}

		return l
	*/
}

func isFreeSpace(l []int, index int) bool {
	return l[index] == -1
}

func isInBounds(l []int, index int) bool {
	return 0 <= index && index < len(l)
}

func findLengthOfStrip(l []int, id int, index int) (int, int) {
	size := 1

	for isInBounds(l, index-size) && l[index-size] == id {
		size++
	}

	return size, index - (size - 1)
}

func findLengthOfStripR(l []int, id int, index int) (int, int) {
	size := 1
	for isInBounds(l, index+size) && l[index+size] == id {
		size++
	}

	return size, index + (size - 1)
}

func findFreeSpace(l []int, size int) int {
	i := 0
	for isInBounds(l, i) {
		if l[i] == -1 {
			length, _ := findLengthOfStripR(l, -1, i)
			//fmt.Println("at index", i, "found free space of size", length)

			if length >= size {
				return i
			} else {
				i += length
			}
		} else {
			i++
		}
	}

	return -1
}

package main

import (
	"adventofcode/util"
	"fmt"
	"log"
)

func main() {
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data []int) int {
	diskMap := expandDiskMap(data)
	defragDisk(diskMap)
	return scoreDisk(diskMap)
}
func part1(data []int) int {
	diskMap := expandDiskMap(data)
	compressDisk(diskMap)
	return scoreDisk(diskMap)
}
func scoreDisk(data []int) int {
	result := 0

	for k, v := range data {
		if v == -1 {
			continue
		}

		result += k * v
	}

	return result
}
func defragDisk(data []int) {
	for i := len(data) - 1; i >= 0; {
		lastFileStart, length := getLastFile(data, i)
		if lastFileStart == -1 {
			return
		}

		firstFreeBlockSpan := getFirstFreeBlockSpan(data, length)
		if firstFreeBlockSpan != -1 && firstFreeBlockSpan < lastFileStart {
			moveFile(data, lastFileStart, length, firstFreeBlockSpan)
		}
		i = lastFileStart - 1
	}
}
func compressDisk(data []int) {
	for {
		lastDataBlock := getLastDataBlock(data)
		firstFreeIndex := getFirstFreeIndex(data)
		if firstFreeIndex > lastDataBlock {
			return
		}

		dataValue := data[lastDataBlock]
		data[firstFreeIndex] = dataValue
		data[lastDataBlock] = -1
	}
}
func getLastFile(data []int, start int) (int, int) {
	fileIndex := -1
	fileStartIndex := -1

	for i := start; i >= 0; i-- {
		if data[i] == -1 {
			if fileIndex == -1 {
				continue
			} else {
				return i + 1, fileStartIndex - i
			}
		}

		if fileIndex == -1 {
			fileIndex = data[i]
			fileStartIndex = i
		} else if data[i] != fileIndex {
			return i + 1, fileStartIndex - i
		}
	}

	return -1, -1
}
func getFirstFreeBlockSpan(data []int, length int) int {
	limit := len(data) - length
	for i := range limit {
		if data[i] != -1 {
			continue
		}

		if isBlockSpanFree(data, i, length) {
			return i
		}
	}

	return -1
}
func isBlockSpanFree(data []int, start, length int) bool {
	for i := range length {
		if data[start+i] != -1 {
			return false
		}
	}

	return true
}
func moveFile(data []int, start, length, target int) {
	for i := range length {
		data[target+i] = data[start+i]
		data[start+i] = -1
	}
}
func getLastDataBlock(data []int) int {
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != -1 {
			return i
		}
	}

	log.Fatalln("failed to find last data block")
	return 0
}
func getFirstFreeIndex(data []int) int {
	for i := range len(data) {
		if data[i] == -1 {
			return i
		}
	}
	log.Fatalln("failed to find first free index")
	return -1
}
func expandDiskMap(data []int) []int {
	result := []int{}

	currentIndex := 0

	for i := range len(data) {
		length := data[i]

		if i%2 == 0 {
			for range length {
				result = append(result, currentIndex)
			}
			currentIndex++
		} else {
			for range length {
				result = append(result, -1)
			}
		}
	}

	return result
}
func diskMapToString(data []int) string {
	result := ""

	for _, v := range data {
		if v != -1 {
			result += fmt.Sprintf("%d", v)
		} else {
			result += "."
		}
	}

	return result
}
func getData(filename string) []int {
	result := []int{}

	lines := util.GetFileLines(filename)

	if len(lines) != 1 {
		log.Fatalln("Expected file with single line!")
	}

	for _, r := range lines[0] {
		result = append(result, int(r)-int('0'))
	}

	return result
}

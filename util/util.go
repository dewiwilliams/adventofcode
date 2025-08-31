package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []string{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func MapStringToArray(source string, mapping map[rune]int) []int {
	result := make([]int, len(source))

	for i, r := range source {
		result[i] = mapping[r]
	}

	return result
}
func ParseIntegerArray(source, separator string) []int {
	result := []int{}

	parts := strings.Split(source, separator)
	for _, p := range parts {
		val, _ := strconv.ParseInt(p, 10, 32)
		result = append(result, int(val))
	}

	return result
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}
func JoinIntArray(values []int, delim string) string {
	stringValues := []string{}
	for _, v := range values {
		stringValues = append(stringValues, strconv.Itoa(v))
	}
	return strings.Join(stringValues, delim)
}

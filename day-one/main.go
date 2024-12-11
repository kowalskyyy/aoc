package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	return string(b[:])
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func SplitList(list []int) (left []int, right []int) {
	for i, num := range list {
		if i%2 == 0 {
			left = append(left, num)

		} else {
			right = append(right, num)
		}
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func calculateDifference(left []int, right []int) (sum int) {
	for i, val := range left {
		sum = sum + absDiffInt(val, right[i])
	}
	return sum
}

func main() {
	input := readInput()
	// tf := "1\n2\n3\n4\n5\n6"
	ints, _ := ReadInts(strings.NewReader(input))
	left, right := SplitList(ints)
	difference := calculateDifference(left, right)
	fmt.Println("result: ", difference)
}

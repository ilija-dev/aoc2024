package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readFileAndSplitIntoArrays(input string) ([]int, []int) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var first, second []int
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) < 2 {
			fmt.Println("input does not contain parts")
		}
		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		first = append(first, firstNum)
		secondNum, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		second = append(second, secondNum)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scanning")
	}
	return first, second
}

func main() {
	first, second := readFileAndSplitIntoArrays("../input.txt")
	// Sort the arrays
	sort.Ints(first)
	sort.Ints(second)

	// Compare oparation
	var sum int
	for i := 0; i < len(first); i++ {
		// Check if difference is positive number
		sum += abs(first[i] - second[i])
	}
	fmt.Println(sum)
}

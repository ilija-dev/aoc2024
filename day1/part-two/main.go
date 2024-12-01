package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	frequencyMap := make(map[int]int)
	var first []int

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 2 {
			continue
		}

		firstNum, _ := strconv.Atoi(parts[0])
		secondNum, _ := strconv.Atoi(parts[1])

		first = append(first, firstNum)
		frequencyMap[secondNum]++
	}

	var sum int
	for _, num := range first {
		sum += num * frequencyMap[num]
	}

	fmt.Println(sum)
}

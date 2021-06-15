package main

import (
	"fmt"
	"strconv"
)

func HandleInfinity() {
	start := argMap["start"]
	startInt, _ := strconv.Atoi(start)

	if startInt == 0 {
		fmt.Println("--start=1 or --start=3 or --start=9")
		return
	}

	DoubleStart(startInt)
}

func DoubleStart(start int) {
	for i := 0; i < 10; i++ {
		fmt.Println("start", start)
		lines := []string{}
		lines = AsciiByteAddition(lines, fmt.Sprintf("%d", start))
		//val := lines[len(lines)-1]
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println("")
		start = start * 2
	}
}

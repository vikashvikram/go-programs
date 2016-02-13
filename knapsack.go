package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(v1, v2 int64) int64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func main() {
	f, err := os.Open("knapsack_big.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	twoD := make([][]int64, 10000)
	var i int64 = 0
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		val, _ := strconv.ParseInt(text[0], 0, 64)
		weight, _ := strconv.ParseInt(text[1], 0, 64)
		twoD[i] = []int64{val, weight}
		i++
	}

	//fmt.Println(twoD[0][0])
	new2D := make([][]int64, twoD[0][1]+1)

	for y := range new2D {
		new2D[y] = make([]int64, twoD[0][0]+1)
	}

	//fmt.Println(new2D[0][0])

	for x := int64(1); x <= twoD[0][1]; x++ {
		for y := int64(0); y <= twoD[0][0]; y++ {
			if y < twoD[x][1] {
				new2D[x][y] = new2D[x-1][y]
			} else {
				val1 := new2D[x-1][y]
				val2 := new2D[x-1][y-twoD[x][1]] + twoD[x][0]
				new2D[x][y] = max(val1, val2)
			}
		}
	}

	fmt.Println(new2D[twoD[0][1]][twoD[0][0]])

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

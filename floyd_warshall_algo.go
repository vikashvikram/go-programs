package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func min(v1, v2 int64) int64 {
	if v1 > v2 {
		return v2
	}
	return v1
}
func main() {
	f, err := os.Open("g3.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	dict := make(map[int64]map[int64]int64)
	nodes := int64(0)
	edges := nodes
	var i int64 = 0
	for scanner.Scan() {
		if i > 0 {
			text := strings.Split(scanner.Text(), " ")
			node1, _ := strconv.ParseInt(text[0], 0, 64)
			node2, _ := strconv.ParseInt(text[1], 0, 64)
			weight, _ := strconv.ParseInt(text[2], 0, 64)
			if dict[node1] == nil {
				dict[node1] = map[int64]int64{node2: weight}
			} else {
				dict[node1][node2] = weight
			}
		} else {
			text := strings.Split(scanner.Text(), " ")
			nodes, _ = strconv.ParseInt(text[0], 0, 64)
			edges, _ = strconv.ParseInt(text[1], 0, 64)
			dict[0] = map[int64]int64{nodes: edges}
		}
		i++
	}

	max := int64(10000)

	var matrix [1001][1001][1001]int64

	for i := int64(1); i <= 1000; i++ {
		for j := int64(1); j <= 1000; j++ {
			if i == j {
				matrix[i][j][0] = int64(0)
			} else if val, ok := dict[i][j]; ok {
				matrix[i][j][0] = val
			} else {
				matrix[i][j][0] = max
			}
		}
	}
	fmt.Println(matrix[1][1][0])

	minimum := max

	count := 1000

	for z := 1; z <= count; z++ {
		for x := 1; x <= count; x++ {
			for y := 1; y <= count; y++ {
				a := matrix[x][y][z-1]
				b := matrix[x][z][z-1]
				c := matrix[z][y][z-1]
				d := b + c
				matrix[x][y][z] = min(a, d)
				if x != y {
					if z == count {
						if minimum > matrix[x][y][z] {
							minimum = matrix[x][y][z]
							fmt.Print("minimum ------** ")
							fmt.Println(minimum)
						}
					}
				}
			}
		}
	}

	for i := int64(1); i <= 1000; i++ {
		if matrix[i][i][1000] < 0 {
			fmt.Println("there is a negative cycle")
		}
	}

	fmt.Print("minimum ------ ")
	fmt.Println(minimum)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type union_find struct {
	parent []int64
	count  []int64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initialize(node_counts int64) union_find {
	uf := union_find{parent: make([]int64, node_counts+1), count: make([]int64, node_counts+1)}
	for i := int64(0); i < node_counts+1; i++ {
		uf.parent[i] = i
		uf.count[i] = 1
	}
	return uf
}

func find(uf union_find, num int64) int64 {
	inp := num
	for {
		num = uf.parent[inp]
		if num == inp {
			break
		} else {
			inp = num
		}
	}
	return inp
}

func union(uf union_find, val1 int64, val2 int64) int64 {
	leader1 := find(uf, val1)
	leader2 := find(uf, val2)
	ret := leader2
	if leader1 != leader2 {
		if uf.count[leader1] > uf.count[leader2] {
			uf.parent[leader2] = leader1
			uf.count[leader1] += uf.count[leader2]
			ret = leader1
		} else {
			uf.parent[leader1] = leader2
			uf.count[leader2] += uf.count[leader1]
		}
	}
	return ret
}

func main() {
	f, err := os.Open("clustering1.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	dict := make(map[int64][]int64)
	count := int64(0)
	clusters := count
	var i int64 = 0
	for scanner.Scan() {
		if i > 0 {
			text := strings.Split(scanner.Text(), " ")
			node1, _ := strconv.ParseInt(text[0], 0, 64)
			node2, _ := strconv.ParseInt(text[1], 0, 64)
			weight, _ := strconv.ParseInt(text[2], 0, 64)
			dict[weight] = append(dict[weight], node1, node2)
		} else {
			count, _ = strconv.ParseInt(scanner.Text(), 0, 64)
			clusters = count
		}
		i++
	}
	var keys []int
	for k := range dict {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	uf := initialize(count)
	for _, k := range keys {
		j := len(dict[int64(k)])
		for x := 0; x < j; x += 2 {
			val1 := dict[int64(k)][int64(x)]
			val2 := dict[int64(k)][int64(x+1)]
			if find(uf, val1) != find(uf, val2) {
				if clusters > 4 {
					union(uf, val1, val2)
					clusters -= 1
				} else {
					fmt.Println("Key:", k)
					os.Exit(3)
				}
			}
		}
	}

}

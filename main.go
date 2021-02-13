package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main () {
	// var c Cache
	// var evPolicy EvictionPolicyTypes
	// evPolicy = LRU
	// c.InitCache(5, evPolicy)
	// c.Set("a", "b")
	// fmt.Println(c.Get("a"))

	// fmt.Println(c)
	list := list.New()
	list.PushBack("A")
	//ele := list.Front()

	m := [][]int{
		{1, 3, 0},
		{5, 2, 0},
	}

	sort.Slice(m, func (i, j int) bool {
		for x := range m[i] {
			if m[i][x] == m[j][x] {
				continue
			}
			return m[i][x] < m[j][x]
		}
		return false
	})
	//sort.Ints(arr)

	fmt.Println(m)
}

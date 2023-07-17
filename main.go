package main

import "fmt"

func main() {
	fullList := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	listUnderFive := make([]int, 0)

	for idc := 0; idc < len(fullList); idc++ {
		if fullList[idc] < 5 {
			listUnderFive = append(listUnderFive, fullList[idc])
		}
	}

	fmt.Println(fullList)
	fmt.Println(listUnderFive)
}

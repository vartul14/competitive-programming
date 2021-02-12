package main

import (
	"fmt"
)

func main () {
	A := []int{1, 2 ,3, 5}
	B := 2

	ans := solve(A, B)
	fmt.Println(ans)
}

func solve(A []int, B int) int {
	ans := 0
	prefixSum := make([]int, len(A))
	suffixSum := make([]int, len(A))
	
    for i, val := range A {
		lastInd := len(A) - i - 1
		if i == 0 {
			prefixSum[i] = val
			suffixSum[lastInd] = A[lastInd]
		} else {
			prefixSum[i] = prefixSum[i - 1] + val
			suffixSum[lastInd] = suffixSum[lastInd + 1] + A[lastInd]
		}
		
	}

	currSum := 0
	for i := 0; i < B; i++ { 
		if i == 0 {
			currSum = suffixSum[len(A) - 1 - B]
		}
		currSum = prefixSum[i] + suffixSum[len(A) - i - 1]
		if currSum > ans {
			ans = currSum
		}

	}

	for i, j := 0, 0; i < 5 && j < 2; i, j = i + 1, j + 1 {
		fmt.Println(i, j)
	}
	
	for i, j := 0, 0; i < 2 && j < 1; i, j = i + 1, j + 1 {
		fmt.Println(i, j)
	}


	return ans

}
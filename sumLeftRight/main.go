package main

import "fmt"

func isBalanced(arr []int, sumArr int) string {
	sumLeft := 0
	sumRight := sumArr - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if sumLeft == sumRight {
			return "YES"
		}

		sumLeft += arr[i]
		sumRight -= arr[i+1]
	}

	return "NO"
}

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)

	sumRight := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sumRight += a[i]
	}

	fmt.Println(isBalanced(a, sumRight))
}

package main

import "fmt"

func main() {
	var b = []int{20, 40, 56, 83, 89}
	var ans float64
	ans = avgVal(b, 5)
	fmt.Printf("The avg is %d\n", ans)
}

func avgVal(arr []int, size int) float64 {
	var i, sum int
	var ans float64
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	ans = float64(sum / size)
	return ans

}

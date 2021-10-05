package fourth

import (
	"fmt"
	"math"
)

func ArrOperations() {
	arr := [6]int{17, 18, 5, 4, 6, 1}
	if len(arr) <= int(math.Pow(10, 4)) {
		for _, val := range arr {
			if val > int(math.Pow(10, 5)) {
				return
			}
		}
		fmt.Println("In:", arr)
		maxRight := arr[len(arr)-1]
		arr[len(arr)-1] = -1
		for i := len(arr) - 2; i >= 0; i-- {
			t := arr[i]
			arr[i] = maxRight
			if t > maxRight {
				maxRight = t
			}
		}
	}
	fmt.Println("Out:", arr)
}

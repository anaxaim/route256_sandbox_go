package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPricesFrequency(pricesArray []int) map[int]int {
	freq := make(map[int]int)
	for _, pr := range pricesArray {
		freq[pr] = freq[pr] + 1
	}
	return freq
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var items int
		fmt.Fscan(in, &items)

		pricesArray := make([]int, items)

		for i := 0; i < items; i++ {
			var nextPrice int
			fmt.Fscan(in, &nextPrice)
			pricesArray[i] = nextPrice
		}

		freq := getPricesFrequency(pricesArray)

		var discountSum int
		for k, v := range freq {
			var s = v - v/3
			discountSum += k * s
		}
		fmt.Fprintln(out, discountSum)
	}
}

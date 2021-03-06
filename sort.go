package main

import "fmt"

func Merge(trans []Transaction, low int, mid int, high int) {
	i := low
	j := mid + 1
	t := trans[low:high]

	//log
	fmt.Println("high: %d", high)

	for k := low; k <= high; k++ {
		fmt.Println("k: %d", k)
		fmt.Println("j: %d", j)
		fmt.Println("i: %d", i)
		if (j <= high) && (i > mid || trans[i].Compare(t[j]) > 0) {
			trans[k] = t[j]
			j++
		} else {
			trans[k] = t[i]
			i++
		}
	}
}

func Sort(trans []Transaction, low int, high int) {
	//base case
	if high <= low {
		return
	}

	mid := low + (high-low)/2
	Sort(trans, low, mid)
	Sort(trans, mid+1, high)
	Merge(trans, low, mid, high)
}

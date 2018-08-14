package main

import (
	"time"

)

func Merge(trans []transactions, low int, high int) {
    i := low
    j := mid + 1


}

func Sort(trans []transactions, low int, high int) {
    if (high <= low) return;
    mid := low + (high - low) / 2
    Sort(trans, low, mid)
    Sort(trans, mid + 1, high)
    Merge(trans, low, mid, high)
}

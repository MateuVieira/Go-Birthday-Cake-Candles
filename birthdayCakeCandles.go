package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	sort.Slice(candles, func(i, j int) bool { return candles[i] < candles[j] })

	lastPosition := int32(len(candles) - 1)
	var tallestCandle int32 = candles[lastPosition]
	var numberOfCandles int32 = 0

	for i := lastPosition; i >= 0; i-- {
		if tallestCandle == candles[i] {
			numberOfCandles++
		}

		if (i - 1) > 0 {
			if tallestCandle > candles[i-1] {
				break
			}
		}
	}

	return numberOfCandles
}

func main() {
	candles := []int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1}

	result := birthdayCakeCandles(candles)

	fmt.Printf("Result: %d\n", result)
}

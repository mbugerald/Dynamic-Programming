package main

import "log"

func main() {
	// memo := map[int]bool{}
	log.Println(canSum(300, []int{7, 14}))
	// log.Println(canSumMemoized(300, []int{7, 14}, memo))
}

func canSum(targetSum int, numbers []int) bool {

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for i := range numbers {

		remainder := targetSum - numbers[i]
		remainderResult := canSum(remainder, numbers) // recursive call.

		if remainderResult == true {
			return true
		}

	}
	return false
}

func canSumMemoized(targetSum int, numbers []int, memo map[int]bool) bool {

	if _, ok := memo[targetSum]; ok {
		return memo[targetSum]
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for i := range numbers {

		remainder := targetSum - numbers[i]
		remainderResult := canSumMemoized(remainder, numbers, memo) // recursive call.

		if remainderResult == true {
			memo[targetSum] = remainderResult
			return memo[targetSum]
		}

	}

	memo[targetSum] = false
	return false
}

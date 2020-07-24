package main

// Subarray Sum Equals K
// Using a map
// O(n) time and space complexity

import "fmt"

func subarraySum(nums []int, k int) int {
	var count, sum int = 0, 0
	m := make(map[int]int)
	// zero sums happens at least once
	m[0] = 1
	// range over list
	for index, _ := range nums {
		// sum up the whole list
		sum += nums[index]
		// check if the current sum of the list, minus k (the searched for solution)
		// is in the map
		if _, ok := m[sum-k]; ok {
			// the reason (sum-k) is mapped, is because a (sum-k) can happen more than once
			// if so increment the count!
			count += m[sum-k]
		}
		// map the sum. Plus the sum of the number of times it happens
		// 1+1 = 2
		// 2:1
		// 0+2 = 2
		// 2:2
		m[sum] = m[sum] + 1
		fmt.Println(m)
	}
	return count
}

func main() {
	in := []int{3, 4, 7, 2, -3, 1, 4, 2}
	num := 7
	out := subarraySum(in, num)
	fmt.Println(out)
}

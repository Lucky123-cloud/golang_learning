package main

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		compliment := target - num

		j, ok := seen[compliment]
		if ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

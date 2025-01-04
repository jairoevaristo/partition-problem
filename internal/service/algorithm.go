package service

import "sort"

type State struct {
	index  int
	target int
}

func sumInt(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}

func hasArrSumImpar(value int) bool {
	return value%2 != 0
}

func hasPartition(index int, currentSum int, targetSum int, arr []int) bool {
	if currentSum == targetSum {
		return true
	}
	if index >= len(arr) || currentSum > targetSum {
		return false
	}

	return hasPartition(index+1, currentSum+arr[index], targetSum, arr) || hasPartition(index+1, currentSum, targetSum, arr)
}

func FindPartitionDynamic(arr []int) bool {
	totalSum := sumInt(arr)

	if hasArrSumImpar(totalSum) {
		return false
	}

	rows := totalSum/2 + 1
	columns := len(arr) + 1

	part := make([][]bool, rows)
	for i := range part {
		part[i] = make([]bool, columns)
	}

	for i := range part[0] {
		part[0][i] = true
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			part[i][j] = part[i][j-1]
			if i >= arr[j-1] {
				part[i][j] = part[i][j] || part[i-arr[j-1]][j-1]
			}
		}
	}

	return part[rows-1][columns-1]
}

func FindPartitionBruteForce(arr []int) bool {
	totalSum := sumInt(arr)

	if hasArrSumImpar(totalSum) {
		return false
	}

	targetSum := totalSum / 2

	return hasPartition(0, 0, targetSum, arr)
}

func HasPartitionGreedy(arr []int) bool {
	totalSum := sumInt(arr)

	if hasArrSumImpar(totalSum) {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	target := totalSum / 2
	sum1, sum2 := 0, 0

	for _, num := range arr {
		if sum1 <= sum2 {
			sum1 += num
		} else {
			sum2 += num
		}
	}

	return sum1 == target && sum2 == target
}

func HasPartitionDC(arr []int) bool {
	totalSum := sumInt(arr)

	if hasArrSumImpar(totalSum) {
		return false
	}

	target := totalSum / 2

	stack := []State{{index: 0, target: target}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.target == 0 {
			return true
		}

		if current.index >= len(arr) || current.target < 0 {
			continue
		}

		stack = append(stack, State{index: current.index + 1, target: current.target - arr[current.index]})
		stack = append(stack, State{index: current.index + 1, target: current.target})
	}

	return false
}

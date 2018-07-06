package threesum

import (
	"fmt"
	"sort"
)

// 穷举无法通过时间限制
func threeSum1(nums []int) [][]int {
	var data [][]int
	if len(nums) < 3 {
		return data
	}
	b := true
	for _, v := range nums {
		if v != 0 {
			b = false
		}
	}
	if b {
		data = append(data, []int{0, 0, 0})
		return data
	}
	cm := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if _, ok := cm[min(nums[i], nums[j], nums[k])]; ok {
						continue
					}

					cm[min(nums[i], nums[j], nums[k])] = true
					data = append(data, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return data
}

func min(i, j, k int) string {
	var t int
	if i > j {
		t, i = i, j
		j = t
	}
	if i > k {
		t, i = i, k
		k = t
	}
	if j > k {
		t, j = j, k
		k = t
	}
	return fmt.Sprintf("%d%d%d", i, j, k)
}

func threeSum(nums []int) [][]int {
	var data [][]int
	ln := len(nums)
	if ln < 3 {
		return data
	}
	sort.Ints(nums)

	var sum int
	for i := range nums {

		target := -nums[i]
		l, r := i+1, ln-1

		for l < r {
			sum = nums[l] + nums[r]
			switch true {
			case sum < target:
				l++
			case sum > target:
				r--
			default:
				tmp := []int{nums[i], nums[l], nums[r]}

				data = append(data, tmp)

				for l < r && nums[l] == tmp[1] {
					l++
				}
				for l < r && nums[r] == tmp[2] {
					r--
				}
			}
		}
		for i+1 < ln && nums[i+1] == nums[i] {
			i++
		}
	}
	return data
}

func twoSum(nums []int, target int) []int {
	fmt.Println("n:", nums)
	fmt.Println("tar:", target)
	data := make([]int, 0, 2)
	if len(nums) < 2 {
		return data
	}

	check := make(map[int]int)
	for k, num := range nums {
		if j, ok := check[num]; ok {
			data = append(data, k, j)
			return data
		}
		check[target-num] = k
	}
	return data
}

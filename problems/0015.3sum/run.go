package threesum

import "fmt"

func threeSum(nums []int) [][]int {
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

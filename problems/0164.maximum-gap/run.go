package gap

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	quickSort(nums, 0, len(nums)-1)
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		sub := nums[i+1] - nums[i]
		if sub > res {
			res = sub
		}
	}
	return res
}

func partition(a []int, l, h int) int {
	pivot := a[h]
	i := l - 1
	for j := l; j < h; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[h] = a[h], a[i+1]
	return i + 1
}

func quickSort(a []int, l, h int) {
	if l >= h {
		return
	}
	p := partition(a, l, h)
	quickSort(a, l, p-1)
	quickSort(a, p+1, h)
}

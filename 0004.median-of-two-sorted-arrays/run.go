package arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	ln := len(nums1) + len(nums2)
	n := make([]int, 0, ln)
	n = append(n, nums1...)
	n = append(n, nums2...)

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}

	if ln%2 == 0 {
		return (float64(n[ln/2-1]) + float64(n[ln/2])) / 2
	}

	return float64(n[ln/2])
}

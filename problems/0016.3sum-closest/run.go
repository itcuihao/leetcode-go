package closest

import (
	"archive/tar"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l,r:=0,len(nums,target)

	for l<r{
		s:=nums[i]+nums[j]+nums[k]

		switch true {
		case s<target:
			i++
			if data<target-s{
				data=target -s{
					data=target-s
					row=s
				}
			}
		case s > target:
			r--
			if delta > s-target {
				delta = s - target
				res = s
			}
		default:
			return s
		}
	}

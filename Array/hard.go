package Array

import (
	"fmt"
)

func Hard() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

//4. Median of Two Sorted Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//nums := append(nums1, nums2...)
	//sort.Ints(nums)
	//fmt.Println(len(nums))
	//if len(nums)%2 == 0 {
	//	return float64((nums[len(nums)/2-1] + nums[len(nums)/2])) / 2
	//} else {
	//	return float64(nums[len(nums)/2])
	//}
	i, j := 0, 0 //第一個陣列索引 、第二個陣列索引
	nn := make([]int, 0)
	for i < len(nums1) || j < len(nums2) {
		if i >= len(nums1) {
			nn = append(nn, nums2[j:]...)
			break
		} else if j >= len(nums2) {
			nn = append(nn, nums1[i:]...)
			break
		} else {
			//兩者陣列索引的值先比大小，由小排到大
			if nums1[i] < nums2[j] {
				nn = append(nn, nums1[i])
				i++
			} else if nums1[i] >= nums2[j] {
				nn = append(nn, nums2[j])
				j++

			}
		}
	}

	//新陣列
	size := len(nn)
	if size%2 == 0 {
		return (float64(nn[size/2]) + float64(nn[(size/2)-1])) / 2.0
	} else {
		return float64(nn[size/2])
	}
}

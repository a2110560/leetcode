package Array

import (
	"fmt"
	"math"
	"sort"
)

func Medium() {
	//fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(reverse(21474836))
	fmt.Println(longestPalindrome("babad"))
}

//5. Longest Palindromic Substring 要在O(n2)以下
//最長回文
//思路 : 從中間開始 假設最中間是第i個元素 如果i-1與i+1不是一樣的，那i+2與i+1就算一樣也不是回文
func longestPalindrome(s string) string {
	res := ""
	reslen := 0
	for i := 0; i < len(s); i++ {
		//長度奇數時
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			//判斷是不是現在最長
			if r-l+1 > reslen {
				res = s[l : r+1]
				fmt.Println(res)
				reslen = r - l + 1
			}
			l--
			r++
		}
		//長度偶數時
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > reslen {
				res = s[l : r+1]
				fmt.Println(res)
				reslen = r - l + 1
			}
			l--
			r++
		}
	}
	if len(res) == 0 {
		res = s[0:1]
	}
	return res
}

//11. Container With Most Water
func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		h := height[l]
		if height[l] > height[r] {
			h = height[r]
		}
		if max < h*(r-l) {
			max = h * (r - l)
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	//超過時間
	//for i := 0; i < len(height); i++ {
	//	for j := i + 1; j < len(height); j++ {
	//		if height[i] > height[j] {
	//			//左邊比右邊高
	//			area := height[j] * (j - i)
	//			if area > max {
	//				max = area
	//			}
	//		}else {
	//			area := height[i] * (j - i)
	//			if area > max {
	//				max = area
	//			}
	//		}
	//	}
	//
	//}
	return max
}

//15. 3Sum
//陣列裡面三個數+起來=0 ， 每個陣列的元素都只能用一次
//陣列長度為6
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
// index=0 012 013 014 015   023 024 025   034 035   045
// index=1 123 124 125 134 135 145
// index=2 234 235 245
// index=3 345
func threeSum(nums []int) [][]int {
	//沒有檢查重複的
	//sort.Ints(nums)
	//var res [][]int
	//var target []int
	//i := 0
	//for i < len(nums)/2+1 {
	//	start, end := i+1, i+2
	//	b := nums[i]
	//	for start <= len(nums)/2+1 && end < len(nums) {
	//		if nums[start]+nums[end] == -b {
	//			target = []int{b, nums[start], nums[end]}
	//			res = append(res, target)
	//		}
	//		if end == len(nums)-1 {
	//			start++
	//			end = i + start
	//		}
	//		end++
	//	}
	//	i++
	//}
	//return res

	//1.先排序
	//2.取第一個num 後面用two sum，雙指標
	//3.假設後面兩組和+起來>0，右邊--
	//4.假設後面兩組和+起來<0，左邊++
	//[-2,-2,0,0,2,2]
	//loop1 : -2
	var res []int
	var ans [][]int
	sort.Ints(nums)
	for index, item := range nums {
		//先排除掉重複
		if index > 0 && item == nums[index-1] {
			continue
		}
		//two pointer
		l, r := index+1, len(nums)-1
		for l < r {
			threesum := item + nums[l] + nums[r]
			if threesum > 0 {
				r--
			} else if threesum < 0 {
				l++
			} else {
				//找到解
				res = []int{item, nums[l], nums[r]}
				ans = append(ans, res)
				l += 1
				//如果下一個又重複
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return ans
}

//16. 3Sum Closest
func threeSumClosest(nums []int, target int) int {
	closet := nums[0] + nums[1] + nums[2]
	diff := math.Abs(float64(closet - target))
	sort.Ints(nums)
	for index, item := range nums {
		if index > 0 && item == nums[index-1] {
			continue
		}
		l, r := index+1, len(nums)-1
		for l < r {
			threesum := item + nums[l] + nums[r]
			newdiff := math.Abs(float64(threesum - target))
			if diff > newdiff {
				diff = newdiff
				closet = threesum
			}
			if threesum < target {
				l++
			} else {
				r--
			}
		}
	}
	return closet
}

//167. Two Sum II - Input Array Is Sorted
func twoSum(numbers []int, target int) []int {
	//map寫法
	//c := make(map[int]int)
	//for index, item := range numbers {
	//	c[item] = index
	//}
	//for i := 0; i < len(numbers); i++ {
	//	a := numbers[i]
	//	b := target - a
	//	if c[b] > 0 {
	//		return []int{i + 1, c[b] + 1}
	//	}
	//}
	//return []int{0}

	var o []int
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return append(o, l+1, r+1)
		} else if l < target {
			l++
		} else {
			r--
		}
	}
	return o
}

//7. Reverse Integer
func reverse(x int) int {
	ans := 0
	for ; x != 0; x /= 10 {
		ans = ans*10 + x%10
	}
	if ans < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	} else {
		return ans
	}
}

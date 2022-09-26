package Array

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Easy() {
	//fmt.Println(isPalindrome(101))
	//fmt.Println(romanToInt("IIIIIVIC"))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(sortArrayByParity([]int{1, 2, 6, 9}))
	//fmt.Println(arrayPairSum([]int{1, 2, 6, 9}))
	//fmt.Println(removeDuplicates([]int{1, 2, 2, 2, 3, 9}))
	//fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	//fmt.Println(plusOne([]int{3, 6, 1, 8}))
	//fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	//fmt.Println(isValid("()()(){[()[]]}"))
	//fmt.Println(getRow(5))
	//fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	//fmt.Println(containsDuplicate([]int{4, 1, 2, 3, 5}))
	//fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	//fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	//fmt.Println(moveZeroes([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

// 9. Palindrome Number
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else {
		count := 0
		for x > count {
			count = count*10 + x%10
			x /= 10
		}
		return x == count || x == count/10
	}
}

//13. Roman to Integer
func romanToInt(s string) int {
	var m = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			{
				if i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
					result -= 1
				} else {
					result += 1
				}
				continue
			}
		case 'X':
			{
				if i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
					result -= 10
				} else {
					result += 10
				}
				continue
			}
		case 'C':
			{
				if i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
					result -= 100
				} else {
					result += 100
				}
				continue
			}
		}
		result += m[s[i]]
	}
	return result
}

//14. Longest Common Prefix
func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return ""
	} else {
		//後面元素都與第一個比，從第一個字開始
		for index, char := range strs[0] {
			for _, str := range strs {
				if index > len(str)-1 {
					return prefix
				}
				if rune(str[index]) != char {
					return prefix
				}
			}
			prefix += string(char)
		}
		return prefix
	}
}

//905. Sort Array By Parity
func sortArrayByParity(nums []int) []int {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j]%2 == 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
		}
	}
	return nums
}

//561. Array Partition
func arrayPairSum(nums []int) int {
	var res int
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

//26. Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) []int {
	inputlen := len(nums)
	count := 0
	nums = append(nums, nums[0])
	for i := 0; i < inputlen; i++ {
		if nums[i] == nums[i+1] {
			continue
		} else {
			nums = append(nums, nums[i+1])
			count++
		}
	}
	nums = nums[inputlen : len(nums)-1]
	return nums
}

//27. Remove Element
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

//66. Plus One
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	//防個位數
	digits = append([]int{1}, digits...)
	return digits
}

//88. Merge Sorted Array
func merge(nums1 []int, m int, nums2 []int, n int) {
	//for i := 0; i < m; i++ {
	//	nums1 = append(nums1, nums1[i])
	//}
	//nums1 = nums1[:m]
	//for i := 0; i < n; i++ {
	//	nums2 = append(nums2, nums2[i])
	//}
	//nums2 = nums2[:n]
	//nums1 = append(nums1, nums2...)
	//sort.Ints(nums1)

	// merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	//i是第一個陣列要取的個數 j是第二個陣列要取的個數
	//k為新陣列的總個數(這裡是沿用舊陣列nums1)
	//i , j , k
	//2 , 2 , 5 ==>一開始
	//2 , 1 , 4 ==>第一次for
	//2 , 0 , 3 ==>第二次for
	//1 , 0 , 2 ==>第三次for
	//1 , -1, 1 ==>第四次for
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 && k >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
		fmt.Println(i, j, k)
		fmt.Println(nums1, nums2)
	}
	return
}

//20. Valid Parentheses
func isValid(s string) bool {
	order := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	//{[(()())]}

	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			lastIndex := len(stack) - 1
			if lastIndex == -1 {
				return false
			}
			lastOpenParenthese := stack[lastIndex]
			if v == order[lastOpenParenthese] {
				stack = stack[:lastIndex]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

//118. Pascal's Triangle
func generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int, i+1) //size of each row will be equal to the row number
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(arr[i]); j++ {
			if j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 1 //每一個元素的陣列最前最後
			} else {
				//i=3 j=1 arr[3][1] = arr[2][0]+arr[2][1]
				//i=3 j=2 arr[3][2] = arr[2][1]+arr[2][2]
				arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			}
		}
	}
	return arr
}

//119. Pascal's Triangle II
func getRow(rowIndex int) []int {
	arr := make([]int, rowIndex+1)
	for i := range arr {
		arr[i] = 1 //全部都塞一
	}
	if rowIndex == 0 {
		arr[0] = 1
		return arr
	} else {
		for i := 0; i < rowIndex; i++ {
			if i == 0 || i == len(arr)-1 {
				arr[i] = 1
			} else {
				for j := i; j >= 1; j-- {
					arr[j] += arr[j-1]
				}
			}
		}
	}
	return arr
}

//121. Best Time to Buy and Sell Stock
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max, lowest := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[lowest] > prices[i] {
			lowest = i
		} else if curr := prices[i] - prices[lowest]; curr > max {
			max = curr
		}
	}
	return max
}

//136. Single Number
func singleNumber(nums []int) int {
	//A XOR A = 0
	//A XOR B = 1
	//B XOR B = 0
	res := 0
	if len(nums) == 1 {
		return nums[0]
	} else {
		for i := 0; i < len(nums); i++ {
			res ^= nums[i] //此舉會把所有的元素拿來XOR，兩個數字XOR之後會變0
		}
	}
	return res
}

//169. Majority Element
func majorityElement(nums []int) int {
	//[11,1,1,1,2,2,2,2,11]
	count := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count++
		} else if nums[i] == result {
			count++
		} else {
			count--
		}
	}
	return result
	//sort.Ints(nums)
	//return nums[len(nums)/2]
}

//217. Contains Duplicate
func containsDuplicate(nums []int) bool {
	//[11,1,1,1,2,2,2,2,11]
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]-nums[j] == 0 {
	//			return true
	//		}
	//	}
	//}
	//return false
	tmp := make(map[int]int)
	for _, v := range nums {
		tmp[v]++
		if tmp[v] > 1 {
			return true
		}
	}
	return false
}

//219. Contains Duplicate II
//(除了要檢查有沒有重複以外，假設有重複，索引值的差不能大於k)
func containsNearbyDuplicate(nums []int, k int) bool {
	//前面int是代表數字，後面int為索引值
	tmp := make(map[int]int)
	for index, item := range nums {
		//先看map裡面有沒有存在，第一組重複可以過，後面重複就會被擋下來
		val, exists := tmp[item]
		if exists && int(math.Abs(float64(index-val))) <= k {
			return true
		}
		tmp[item] = index
	}
	return false
}

////228. Summary Ranges(有排小到大，值相異) =>把連一起的寫在一起
/*
輸入： nums = [0,2,3,4,6,8,9]
輸出： ["0","2->4","6","8->9"]
*/
func summaryRanges(nums []int) []string {
	var res []string
	start, last := 0, 0 //端點
	if len(nums) == 0 {
		return []string{}
	} else {
		for i := 1; i < len(nums); i++ {
			if nums[i]-nums[i-1] == 1 {
				//把區段的終點往後移
				last = i
			} else {
				if start == last {
					//代表只有一個
					res = append(res, strconv.Itoa(nums[start]))
				} else {
					//代表有區段 2,3,4
					res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[last]))
				}
				start, last = i, i
			}
		}
	}

	//處理最後一個
	if start == last {
		return append(res, strconv.Itoa(nums[start]))
	} else {
		return append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[last]))
	}

}

//268. Missing Number
func missingNumber(nums []int) int {
	//sort.Ints(nums)
	//m := make(map[int]int)
	//for _, item := range nums {
	//	m[item]++
	//}
	//for i := 0; i <= len(nums); i++ {
	//	if m[i] == 0 {
	//		return i
	//	}
	//}
	//return 0
	var sum int
	var ind int
	for i, v := range nums {
		sum += v
		ind += i + 1
	}
	return ind - sum
}

//69. Sqrt(x)
func mySqrt(x int) int {
	//for i := 0; ; i++ {
	//	// x=   4-9
	//	if x >= i*i && x <= (i+1)*(i+1) {
	//		if x == (i+1)*(i+1) {
	//			return i + 1
	//		} else {
	//			return i
	//		}
	//	}
	//}
	// x=8
	l := 0
	r := x
	for l < r {
		m := (r-l+1)/2 + l
		if m*m > x {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

//70. Climbing Stairs
func climbStairs(n int) int {
	//遞迴版
	//if n==1 {
	//	return 1
	//}else if n==2{
	//	return 2
	//}else{
	//	return climbStairs(n-1)+climbStairs(n-2)
	//}
	if n <= 2 {
		return n
	}
	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := n2
		n2 = n1 + n2
		n1 = tmp
	}
	return n2
}

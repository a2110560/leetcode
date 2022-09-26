package String

import "fmt"

func Easy() {
	fmt.Println(AddBinary("1101", "1011"))
}

// AddBinary 67. Add Binary
func AddBinary(a string, b string) string {
	s := 0
	carry := 0
	res := ""
	la := len(a) - 1 //3
	lb := len(b) - 1 //3
	//思路 : 把每一次最後那位取出來 加上s 如果除2等於1代表要進位，等於0代表不用進位
	for la >= 0 || lb >= 0 || carry != 0 {
		s = carry
		if la >= 0 {
			s += int(a[la] - '0')
			la--
		}
		if lb >= 0 {
			s += int(b[lb] - '0')
			lb--
		}
		carry = s / 2
		res = string(rune(s%2+'0')) + res

	}
	return res
}

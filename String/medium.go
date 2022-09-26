package String

import "fmt"

func Medium() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

//6. Zigzag Conversion
func convert(s string, numRows int) string {
	/*
	   P0      I6        N12
	   A1   L5 S7    I11 G13
	   Y2 A4   H8 R10
	   P3      I9
	*/
	//row=4 len=14
	//     P0    A4    H8       N12
	//     A1 P3 L5 S7 I9  I11  G13
	//     Y2    I6    R10
	// row=3 len=14

	if numRows == 1 {
		return s
	}
	var str []string
	cyclelen := 2*numRows - 2 //跑回循環的長度
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += cyclelen {
			str = append(str, string(s[i+j]))
			fmt.Println("out : ", str)
			if i != 0 && i != numRows-1 && j+cyclelen-i < len(s) {
				str = append(str, string(s[j+cyclelen-i]))
				fmt.Println("in : ", str)
			}
		}
	}
	newstr := ""
	for i := range str {
		newstr += str[i]
	}
	return newstr
}

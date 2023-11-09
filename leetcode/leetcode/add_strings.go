package leetcode

import "fmt"

func AddStrings() {
	fmt.Println(addStrings("11", "123")) // 134
	fmt.Println(addStrings("456", "77")) // 533
	fmt.Println(addStrings("0", "0"))    // 0
}

func addStrings(num1 string, num2 string) string {
	p1 := len(num1) - 1
	p2 := len(num2) - 1
	var carry byte
	var sb []byte

	for p1 >= 0 || p2 >= 0 {
		var b1 byte
		if p1 >= 0 {
			b1 = num1[p1] - '0'
		}
		var b2 byte
		if p2 >= 0 {
			b2 = num2[p2] - '0'
		}
		b3 := b1 + b2 + carry
		if b3 > 9 {
			carry = b3 / 10
			b3 = b3 % 10
		} else {
			carry = 0
		}
		sb = append([]byte{b3 + '0'}, sb...)
		p1--
		p2--
	}

	if carry > 0 {
		sb = append([]byte{carry + '0'}, sb...)
	}
	return string(sb)
}

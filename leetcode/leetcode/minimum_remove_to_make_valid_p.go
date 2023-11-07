package leetcode

import (
	"fmt"
)

func MinRemoveToMakeValid() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de"))
	fmt.Println(minRemoveToMakeValid("ab(c)d"))
	fmt.Println(minRemoveToMakeValid("))(("))
}

func minRemoveToMakeValid(s string) string {
	st := &Stack[int]{}
	ss := []rune(s)
	i := 0

	for i < len(ss) {
		if ss[i] == '(' {
			st.Push(i)
		} else if ss[i] == ')' && st.Len() > 0 {
			st.Pop()
		} else if ss[i] == ')' {
			ss = append(ss[:i], ss[i+1:]...)
			i--
		}
		i++
	}

	for st.Len() > 0 {
		i = st.Pop()
		ss = append(ss[:i], ss[i+1:]...)
	}
	return string(ss)
}

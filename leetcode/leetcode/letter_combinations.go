package leetcode

import (
	"fmt"
)

func LetterCombinations() {
	{
		fmt.Println(letterCombinations("23"))
		// ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	}

	{
		fmt.Println(letterCombinations(""))
		// []
	}

	{
		fmt.Println(letterCombinations("2"))
		// ["a","b","c"]
	}

	{
		fmt.Println(letterCombinations("234"))
		// ["adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"]
	}
}

func letterCombinations(digits string) []string {
	m := map[byte]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	L := len(digits)
	if L == 0 {
		return []string{}
	}

	ans := make([]string, 0, L*3)
	var dfs func([]byte, int)
	dfs = func(s []byte, idx int) {
		if idx >= L {
			ans = append(ans, string(s))
			return
		}

		num := digits[idx] - '0'
		letters := m[num]
		N := len(letters)
		for i := 0; i < N; i++ {
			s = append(s, letters[i])
			dfs(s, idx+1)
			s = s[:len(s)-1]
		}
	}
	dfs([]byte{}, 0)

	return ans
}

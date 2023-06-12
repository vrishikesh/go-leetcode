package star

import (
	"fmt"
	"strings"
)

//	 *
//	* *
// * * *

func Triangle() {
	N := 10
	SSize := N - 1
	S := strings.Repeat(" ", SSize)

	stars := ""
	for i := 0; i < N; i++ {
		stars += "* "
		fmt.Println(S[:SSize] + stars)
		SSize -= 1
	}
}

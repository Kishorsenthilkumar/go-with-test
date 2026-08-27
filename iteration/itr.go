package iteration

import "strings"

const nooftimerepeat = 5

func Repeat(letter string) string {
	var repeated strings.Builder
	for i := 0; i < nooftimerepeat; i++ {
		repeated.WriteString(letter)
	}
	return repeated.String()

}

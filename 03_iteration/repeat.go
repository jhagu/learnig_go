package iteration

import "strings"

func Repeat(character string, times int) string {
	if times == 0 {
		times = 5
	}
	/* for i := 0; i < times; i++ {
		repeated += character
	} */
	return strings.Repeat(character, times)
}

//string utility
package strings


func RuneLen(s string) int {
	return len([]rune(s))
}

func RuneSubString(s string, begin, length int) string {
	rs := []rune(s)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}
package pkg01

/*
翻转字符串
*/
func Reverse(s string) string {
	r := []rune(s)
	len := len(r)

	for i, j := 0, len-1; i < len/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

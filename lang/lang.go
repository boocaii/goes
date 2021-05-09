package lang

import "regexp"

var (
	reHan = regexp.MustCompile(`^\p{Han}+$`)
	reLatin = regexp.MustCompile(`^\p{Latin}+$`)
)



func IsHan(s string) bool {
	return reHan.MatchString(s)
}

func IsLatin(s string) bool {
	return reLatin.MatchString(s)
}

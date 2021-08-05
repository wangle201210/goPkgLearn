package convert

import (
	"regexp"
	"strconv"
)

func Octal() (str string) {
	s := "\346\270\240\351\201\223\357\274\214\346\255\243\345\244\247\345\244\251\346\231\264\346\270\240\351\201\223(zdtq)\347\232\204\345\214\273\347\224\237\344\270\215\347\224\250\346\240\241\351\252\214\351\252\214\350\257\201\347\240\201\357\274\214\347\233\264\346\216\245\347\231\273\351\231\206"
	str = convertOctonaryUtf8(s)
	println(str)
	return
}

func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}

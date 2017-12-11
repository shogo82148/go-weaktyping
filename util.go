package weaktyping

import (
	"strconv"
)

func unquoteBytesIfQuoted(s []byte) []byte {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return s
	}
	s = s[1 : len(s)-1]

	// skip decoding escape sequence
	// we assumed s does not contain espace sequence

	return s
}

func parseUint(s string, base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(s, base, bitSize)
}

func parseInt(s string, base int, bitSize int) (int64, error) {
	return strconv.ParseInt(s, base, bitSize)
}

func parseFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}

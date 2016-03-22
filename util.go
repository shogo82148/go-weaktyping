package weaktyping

func unquoteBytesIfQuoted(s []byte) []byte {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return s
	}
	s = s[1 : len(s)-1]

	// skip decoding escape sequence
	// we assumed s does not contain espace sequence

	return s
}

package weaktyping

import "strconv"

type Int int

func PtrInt(v Int) *Int {
	return &v
}

func (v *Int) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Int) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	if i, err := strconv.ParseInt(s, 10, 0); err != nil {
		return err
	} else {
		*v = Int(i)
	}
	return nil
}

func unquoteBytesIfQuoted(s []byte) []byte {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return s
	}
	s = s[1 : len(s)-1]

	// skip decoding escape sequence
	// we assumed s does not contain espace sequence

	return s
}

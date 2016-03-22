package weaktyping

import "strconv"

type Int struct{ Value int }

func (v Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(v.Value), 10)), nil
}

func (v Int) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(v.Value), 10)), nil
}

func (v *Int) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Int) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		v.Value = 0
		return nil
	}
	if i, err := strconv.ParseInt(s, 10, 0); err != nil {
		return err
	} else {
		v.Value = int(i)
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

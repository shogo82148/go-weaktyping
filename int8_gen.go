package weaktyping

import "strconv"

type Int8 int8

func PtrInt8(v Int8) *Int8 {
	return &v
}

func (v *Int8) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Int8) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	if w, err := strconv.ParseInt(s, 10, 8); err != nil {
		return err
	} else {
		*v = Int8(w)
	}
	return nil
}

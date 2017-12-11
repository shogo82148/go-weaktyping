package weaktyping

import "strconv"

type Int16 int16

func PtrInt16(v Int16) *Int16 {
	return &v
}

func (v *Int16) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Int16) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return err
	}
	*v = Int16(w)
	return nil
}

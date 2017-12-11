package weaktyping

import "strconv"

type Int32 int32

func PtrInt32(v Int32) *Int32 {
	return &v
}

func (v *Int32) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Int32) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	*v = Int32(w)
	return nil
}

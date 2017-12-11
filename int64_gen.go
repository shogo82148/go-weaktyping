package weaktyping

import "strconv"

type Int64 int64

func PtrInt64(v Int64) *Int64 {
	return &v
}

func (v *Int64) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Int64) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*v = Int64(w)
	return nil
}

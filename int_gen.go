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
	w, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return err
	}
	*v = Int(w)
	return nil
}

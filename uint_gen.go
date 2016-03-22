package weaktyping

import "strconv"

type Uint uint

func PtrUint(v Uint) *Uint {
	return &v
}

func (v *Uint) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Uint) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	if w, err := strconv.ParseUint(s, 10, 0); err != nil {
		return err
	} else {
		*v = Uint(w)
	}
	return nil
}

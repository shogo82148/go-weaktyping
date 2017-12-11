package weaktyping

import "strconv"

type Uint8 uint8

func PtrUint8(v Uint8) *Uint8 {
	return &v
}

func (v *Uint8) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Uint8) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return err
	}
	*v = Uint8(w)
	return nil
}

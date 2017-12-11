package weaktyping

import "strconv"

type Uint16 uint16

func PtrUint16(v Uint16) *Uint16 {
	return &v
}

func (v *Uint16) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Uint16) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return err
	}
	*v = Uint16(w)
	return nil
}

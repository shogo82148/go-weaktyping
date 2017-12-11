package weaktyping

import "strconv"

type Uint32 uint32

func PtrUint32(v Uint32) *Uint32 {
	return &v
}

func (v *Uint32) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Uint32) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	*v = Uint32(w)
	return nil
}

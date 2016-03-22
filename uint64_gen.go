package weaktyping

import "strconv"

type Uint64 uint64

func PtrUint64(v Uint64) *Uint64 {
	return &v
}

func (v *Uint64) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Uint64) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	if w, err := strconv.ParseUint(s, 10, 64); err != nil {
		return err
	} else {
		*v = Uint64(w)
	}
	return nil
}
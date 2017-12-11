package weaktyping

// Uint64 is weak typed uint64.
type Uint64 uint64

// PtrUint64 returns the pointer of v.
func PtrUint64(v Uint64) *Uint64 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Uint64) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalJSON implements "encoding".TextUnmarshaler.
func (v *Uint64) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*v = Uint64(w)
	return nil
}

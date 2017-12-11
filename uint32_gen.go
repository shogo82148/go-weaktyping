package weaktyping

// Uint32 is weak typed uint32.
type Uint32 uint32

// PtrUint32 returns the pointer of v.
func PtrUint32(v Uint32) *Uint32 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Uint32) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Uint32) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseUint(s, 10, 32)
	if err != nil {
		return err
	}
	*v = Uint32(w)
	return nil
}

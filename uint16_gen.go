package weaktyping

// Uint16 is weak typed uint16.
type Uint16 uint16

// PtrUint16 returns the pointer of v.
func PtrUint16(v Uint16) *Uint16 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Uint16) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Uint16) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseUint(s, 10, 16)
	if err != nil {
		return err
	}
	*v = Uint16(w)
	return nil
}

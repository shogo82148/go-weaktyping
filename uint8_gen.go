package weaktyping

// Uint8 is weak typed uint8.
type Uint8 uint8

// PtrUint8 returns the pointer of v.
func PtrUint8(v Uint8) *Uint8 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Uint8) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Uint8) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "" || s == "null" {
		*v = 0
		return nil
	}
	w, err := parseUint(s, 10, 8)
	if err != nil {
		return err
	}
	*v = Uint8(w)
	return nil
}

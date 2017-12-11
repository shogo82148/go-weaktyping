package weaktyping

// Uint is weak typed uint.
type Uint uint

// PtrUint returns the pointer of v.
func PtrUint(v Uint) *Uint {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Uint) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Uint) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseUint(s, 10, 0)
	if err != nil {
		return err
	}
	*v = Uint(w)
	return nil
}

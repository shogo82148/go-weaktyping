package weaktyping

// Int16 is weak typed int16.
type Int16 int16

// PtrInt16 returns the pointer of v.
func PtrInt16(v Int16) *Int16 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Int16) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalJSON implements "encoding".TextUnmarshaler.
func (v *Int16) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseInt(s, 10, 16)
	if err != nil {
		return err
	}
	*v = Int16(w)
	return nil
}

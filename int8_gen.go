package weaktyping

// Int8 is weak typed int8.
type Int8 int8

// PtrInt8 returns the pointer of v.
func PtrInt8(v Int8) *Int8 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Int8) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalJSON implements "encoding".TextUnmarshaler.
func (v *Int8) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseInt(s, 10, 8)
	if err != nil {
		return err
	}
	*v = Int8(w)
	return nil
}

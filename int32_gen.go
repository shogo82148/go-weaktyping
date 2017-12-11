package weaktyping

// Int32 is weak typed int32.
type Int32 int32

// PtrInt32 returns the pointer of v.
func PtrInt32(v Int32) *Int32 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Int32) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalJSON implements "encoding".TextUnmarshaler.
func (v *Int32) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseInt(s, 10, 32)
	if err != nil {
		return err
	}
	*v = Int32(w)
	return nil
}

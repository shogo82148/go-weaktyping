package weaktyping

// Int is weak typed int.
type Int int

// PtrInt returns the pointer of v.
func PtrInt(v Int) *Int {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Int) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalJSON implements "encoding".TextUnmarshaler.
func (v *Int) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseInt(s, 10, 0)
	if err != nil {
		return err
	}
	*v = Int(w)
	return nil
}

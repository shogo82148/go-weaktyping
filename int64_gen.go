package weaktyping

// Int64 is weak typed int64.
type Int64 int64

// PtrInt64 returns the pointer of v.
func PtrInt64(v Int64) *Int64 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Int64) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Int64) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*v = Int64(w)
	return nil
}

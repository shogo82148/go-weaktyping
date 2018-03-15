package weaktyping

// Float64 is weak typed float64.
type Float64 float64

// PtrFloat64 returns the pointer of v.
func PtrFloat64(v Float64) *Float64 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Float64) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Float64) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "" || s == "null" {
		*v = 0
		return nil
	}
	w, err := parseFloat(s, 64)
	if err != nil {
		return err
	}
	*v = Float64(w)
	return nil
}

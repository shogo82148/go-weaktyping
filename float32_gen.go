package weaktyping

// Float32 is weak typed float32.
type Float32 float32

// PtrFloat32 returns the pointer of v.
func PtrFloat32(v Float32) *Float32 {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Float32) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Float32) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	w, err := parseFloat(s, 32)
	if err != nil {
		return err
	}
	*v = Float32(w)
	return nil
}

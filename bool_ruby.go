package weaktyping

// RubyBool is weak typed bool.
// false and null are treated as false, others are true.
type RubyBool bool

// PtrRubyBool returns the pointer of v.
func PtrRubyBool(v RubyBool) *RubyBool {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *RubyBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	*v = RubyBool(s != "false" && s != "null")
	return nil
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *RubyBool) UnmarshalText(data []byte) error {
	*v = RubyBool(true)
	return nil
}

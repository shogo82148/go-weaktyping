package weaktyping

type RubyBool bool

func PtrRubyBool(v RubyBool) *RubyBool {
	return &v
}

func (v *RubyBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	*v = RubyBool(s != "false" && s != "null")
	return nil
}

func (v *RubyBool) UnmarshalText(data []byte) error {
	*v = RubyBool(true)
	return nil
}

package weaktyping

import "strconv"

// Bool is weak typed bool.
// false, 0, "", and null are treated as false, others are true.
// This behavior is same as JavaScriptBool.
type Bool bool

// PtrBool returns the pointer of v.
func PtrBool(v Bool) *Bool {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *Bool) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch s {
	case `false`, `0`, `""`, `null`:
		*v = false
	default:
		f, err := strconv.ParseFloat(s, 64)
		*v = Bool(err != nil || f != 0)
	}
	return nil
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *Bool) UnmarshalText(data []byte) error {
	*v = len(data) != 0
	return nil
}

package weaktyping

import "strconv"

// JavaScriptBool is weak typed bool.
// false, 0, "", and null are treated as false, others are true.
type JavaScriptBool bool

// PtrJavaScriptBool returns the pointer of v.
func PtrJavaScriptBool(v JavaScriptBool) *JavaScriptBool {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *JavaScriptBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch s {
	case `false`, `0`, `""`, `null`:
		*v = false
	default:
		f, err := strconv.ParseFloat(s, 64)
		*v = JavaScriptBool(err != nil || f != 0)
	}
	return nil
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *JavaScriptBool) UnmarshalText(data []byte) error {
	*v = JavaScriptBool(len(data) != 0)
	return nil
}

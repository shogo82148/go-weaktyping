package weaktyping

import "strconv"

type JavaScriptBool bool

func PtrJavaScriptBool(v JavaScriptBool) *JavaScriptBool {
	return &v
}

func (v *JavaScriptBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch s {
	case `false`, `0`, `""`, `null`:
		*v = false
	default:
		f, err := strconv.ParseFloat(s, 64)
		*v = err != nil || f != 0
	}
	return nil
}

func (v *JavaScriptBool) UnmarshalText(data []byte) error {
	*v = len(data) != 0
	return nil
}

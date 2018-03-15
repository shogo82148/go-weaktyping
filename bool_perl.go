package weaktyping

import "strconv"

// PerlBool is weak typed bool.
// false, 0, "0", "", and null are treated as false, others are true.
type PerlBool bool

// PtrPerlBool returns the pointer of v.
func PtrPerlBool(v PerlBool) *PerlBool {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *PerlBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch s {
	case `false`, `0`, `""`, `null`, `"0"`:
		*v = false
	default:
		f, err := strconv.ParseFloat(s, 64)
		*v = PerlBool(err != nil || f != 0)
	}
	return nil
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *PerlBool) UnmarshalText(data []byte) error {
	*v = PerlBool(len(data) != 0 && (len(data) != 1 || data[0] != '0'))
	return nil
}

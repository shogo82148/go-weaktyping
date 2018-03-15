package weaktyping

import "strconv"

// PHPBool is weak typed bool.
// false, 0, "", [], and null are treated as false, others are true.
type PHPBool bool

// PtrPHPBool returns the pointer of v.
func PtrPHPBool(v PHPBool) *PHPBool {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *PHPBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch s {
	case `false`, `0`, `""`, `null`, `"0"`, `[]`:
		*v = false
	default:
		f, err := strconv.ParseFloat(s, 64)
		*v = PHPBool(err != nil || f != 0)
	}
	return nil
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *PHPBool) UnmarshalText(data []byte) error {
	*v = PHPBool(len(data) != 0 && (len(data) != 1 || data[0] != '0'))
	return nil
}

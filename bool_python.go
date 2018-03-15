package weaktyping

import "strconv"

// PythonBool is weak typed bool.
// false, 0, "", [], {}, and null are treated as false, others are true.
type PythonBool bool

// PtrPythonBool returns the pointer of v.
func PtrPythonBool(v PythonBool) *PythonBool {
	return &v
}

// UnmarshalJSON implements "encoding/json".Unmarshaler.
func (v *PythonBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch s {
	case `false`, `0`, `""`, `null`, `[]`, `{}`:
		*v = false
	default:
		f, err := strconv.ParseFloat(s, 64)
		*v = PythonBool(err != nil || f != 0)
	}
	return nil
}

// UnmarshalText implements "encoding".TextUnmarshaler.
func (v *PythonBool) UnmarshalText(data []byte) error {
	*v = PythonBool(len(data) != 0)
	return nil
}

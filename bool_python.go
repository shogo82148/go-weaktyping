package weaktyping

import "strconv"

type PythonBool bool

func PtrPythonBool(v PythonBool) *PythonBool {
	return &v
}

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

func (v *PythonBool) UnmarshalText(data []byte) error {
	*v = PythonBool(len(data) != 0)
	return nil
}

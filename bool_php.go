package weaktyping

import "strconv"

type PHPBool bool

func PtrPHPBool(v PHPBool) *PHPBool {
	return &v
}

func (v *PHPBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch s {
	case `false`, `0`, `""`, `null`, `"0"`, `[]`:
		*v = false
	default:
		f, err := strconv.ParseFloat(s, 64)
		*v = err != nil || f != 0
	}
	return nil
}

func (v *PHPBool) UnmarshalText(data []byte) error {
	*v = len(data) != 0 && (len(data) != 1 || data[0] != '0')
	return nil
}

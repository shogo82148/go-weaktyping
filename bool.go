package weaktyping

import "strconv"

type Bool bool

func PtrBool(v Bool) *Bool {
	return &v
}

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

func (v *Bool) UnmarshalText(data []byte) error {
	*v = len(data) != 0
	return nil
}

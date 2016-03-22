package weaktyping

import "strconv"

type Float64 float64

func PtrFloat64(v Float64) *Float64 {
	return &v
}

func (v *Float64) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Float64) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	if w, err := strconv.ParseFloat(s, 64); err != nil {
		return err
	} else {
		*v = Float64(w)
	}
	return nil
}

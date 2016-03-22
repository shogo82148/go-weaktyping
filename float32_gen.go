package weaktyping

import "strconv"

type Float32 float32

func PtrFloat32(v Float32) *Float32 {
	return &v
}

func (v *Float32) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *Float32) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	if w, err := strconv.ParseFloat(s, 32); err != nil {
		return err
	} else {
		*v = Float32(w)
	}
	return nil
}

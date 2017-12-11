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
	w, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*v = Float32(w)
	return nil
}

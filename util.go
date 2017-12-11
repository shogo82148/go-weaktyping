package weaktyping

import (
	"math/big"
	"strconv"
)

func unquoteBytesIfQuoted(s []byte) []byte {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return s
	}
	s = s[1 : len(s)-1]

	// skip decoding escape sequence
	// we assumed s does not contain espace sequence

	return s
}

func parseUint(s string, base int, bitSize int) (uint64, error) {
	// parse as uint.
	if v, err := strconv.ParseUint(s, base, bitSize); err == nil {
		return v, nil
	}

	if bitSize == 0 {
		bitSize = int(strconv.IntSize)
	}

	// parse as float
	maxVal := uint64(1<<uint(bitSize) - 1)
	if bitSize <= 53 {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		if v < 0 || v > float64(maxVal) {
			return 0, &strconv.NumError{
				Func: "ParseUint",
				Num:  s,
				Err:  strconv.ErrRange,
			}
		}
		return uint64(v), nil
	}

	// parse as big.Float
	v, _, err := big.ParseFloat(s, 10, 64, big.ToZero)
	if err != nil {
		return 0, err
	}
	ret, a := v.Uint64()
	if a != big.Exact || ret > maxVal {
		return 0, &strconv.NumError{
			Func: "ParseUint",
			Num:  s,
			Err:  strconv.ErrRange,
		}
	}

	return ret, nil
}

func parseInt(s string, base int, bitSize int) (int64, error) {
	// parse as int.
	if v, err := strconv.ParseInt(s, base, bitSize); err == nil {
		return v, nil
	}

	if bitSize == 0 {
		bitSize = int(strconv.IntSize)
	}

	cutoff := int64(1 << uint(bitSize-1))
	if bitSize <= 53 {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		if v < float64(-cutoff) || v > float64(cutoff-1) {
			return 0, &strconv.NumError{
				Func: "ParseUint",
				Num:  s,
				Err:  strconv.ErrRange,
			}
		}
		return int64(v), nil
	}

	// parse as big.Float
	v, _, err := big.ParseFloat(s, 10, 64, big.ToZero)
	if err != nil {
		return 0, err
	}
	ret, a := v.Int64()
	if a != big.Exact || ret < int64(-cutoff) || ret > int64(cutoff-1) {
		return 0, &strconv.NumError{
			Func: "ParseUint",
			Num:  s,
			Err:  strconv.ErrRange,
		}
	}
	return ret, nil
}

func parseFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}

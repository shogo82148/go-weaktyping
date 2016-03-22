package weaktyping

//go:generate sh generate.sh Int int "strconv.ParseInt(s, 10, 0)"
//go:generate sh generate.sh Int8 int8 "strconv.ParseInt(s, 10, 8)"
//go:generate sh generate.sh Int16 int16 "strconv.ParseInt(s, 10, 16)"
//go:generate sh generate.sh Int32 int32 "strconv.ParseInt(s, 10, 32)"
//go:generate sh generate.sh Int64 int64 "strconv.ParseInt(s, 10, 64)"
//go:generate sh generate.sh Uint uint "strconv.ParseUint(s, 10, 0)"
//go:generate sh generate.sh Uint8 uint8 "strconv.ParseUint(s, 10, 8)"
//go:generate sh generate.sh Uint16 uint16 "strconv.ParseUint(s, 10, 16)"
//go:generate sh generate.sh Uint32 uint32 "strconv.ParseUint(s, 10, 32)"
//go:generate sh generate.sh Uint64 uint64 "strconv.ParseUint(s, 10, 64)"
//go:generate sh generate.sh Float32 float32 "strconv.ParseFloat(s, 32)"
//go:generate sh generate.sh Float64 float64 "strconv.ParseFloat(s, 64)"

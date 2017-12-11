package weaktyping

//go:generate sh generate.sh Int int "parseInt(s, 10, 0)"
//go:generate sh generate.sh Int8 int8 "parseInt(s, 10, 8)"
//go:generate sh generate.sh Int16 int16 "parseInt(s, 10, 16)"
//go:generate sh generate.sh Int32 int32 "parseInt(s, 10, 32)"
//go:generate sh generate.sh Int64 int64 "parseInt(s, 10, 64)"
//go:generate sh generate.sh Uint uint "parseUint(s, 10, 0)"
//go:generate sh generate.sh Uint8 uint8 "parseUint(s, 10, 8)"
//go:generate sh generate.sh Uint16 uint16 "parseUint(s, 10, 16)"
//go:generate sh generate.sh Uint32 uint32 "parseUint(s, 10, 32)"
//go:generate sh generate.sh Uint64 uint64 "parseUint(s, 10, 64)"
//go:generate sh generate.sh Float32 float32 "parseFloat(s, 32)"
//go:generate sh generate.sh Float64 float64 "parseFloat(s, 64)"

package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalInt(t *testing.T) {
	// Special test cases for integer type.
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
		err bool
	}{
		// Accept floating values.
		{`{"foo":123.0}`, new(map[string]Int), &map[string]Int{"foo": 123}, false},
		{`{"foo":"123.0"}`, new(map[string]Int), &map[string]Int{"foo": 123}, false},
		{`{"foo":1.23e+2}`, new(map[string]Int), &map[string]Int{"foo": 123}, false},
		{`{"foo":123.0}`, new(map[string]Uint), &map[string]Uint{"foo": 123}, false},
		{`{"foo":"123.0"}`, new(map[string]Uint), &map[string]Uint{"foo": 123}, false},
		{`{"foo":1.23e+2}`, new(map[string]Uint), &map[string]Uint{"foo": 123}, false},

		// range check for int8.
		{`{"foo":-129}`, new(map[string]Int8), nil, true},
		{`{"foo":-128}`, new(map[string]Int8), &map[string]Int8{"foo": -128}, false},
		{`{"foo":127}`, new(map[string]Int8), &map[string]Int8{"foo": 127}, false},
		{`{"foo":128}`, new(map[string]Int8), nil, true},

		// range check for int16.
		{`{"foo":-32769}`, new(map[string]Int16), nil, true},
		{`{"foo":-32768}`, new(map[string]Int16), &map[string]Int16{"foo": -32768}, false},
		{`{"foo":32767}`, new(map[string]Int16), &map[string]Int16{"foo": 32767}, false},
		{`{"foo":32768}`, new(map[string]Int16), nil, true},

		// range check for int32.
		{`{"foo":-2147483649}`, new(map[string]Int32), nil, true},
		{`{"foo":-2147483648}`, new(map[string]Int32), &map[string]Int32{"foo": -2147483648}, false},
		{`{"foo":2147483647}`, new(map[string]Int32), &map[string]Int32{"foo": 2147483647}, false},
		{`{"foo":2147483648}`, new(map[string]Int32), nil, true},
		{`{"foo":2147483647.0}`, new(map[string]Int32), &map[string]Int32{"foo": 2147483647}, false},
		{`{"foo":2147483648.0}`, new(map[string]Int32), nil, true},

		// range check for int64.
		{`{"foo":-9223372036854775809}`, new(map[string]Int64), nil, true},
		{`{"foo":-9223372036854775808}`, new(map[string]Int64), &map[string]Int64{"foo": -9223372036854775808}, false},
		{`{"foo":-9223372036854775807}`, new(map[string]Int64), &map[string]Int64{"foo": -9223372036854775807}, false},
		{`{"foo":-9223372036854775806}`, new(map[string]Int64), &map[string]Int64{"foo": -9223372036854775806}, false},
		{`{"foo":9223372036854775805}`, new(map[string]Int64), &map[string]Int64{"foo": 9223372036854775805}, false},
		{`{"foo":9223372036854775806}`, new(map[string]Int64), &map[string]Int64{"foo": 9223372036854775806}, false},
		{`{"foo":9223372036854775807}`, new(map[string]Int64), &map[string]Int64{"foo": 9223372036854775807}, false},
		{`{"foo":9223372036854775808}`, new(map[string]Int64), nil, true},
		{`{"foo":-9223372036854775809.0}`, new(map[string]Int64), nil, true},
		{`{"foo":-9223372036854775808.0}`, new(map[string]Int64), &map[string]Int64{"foo": -9223372036854775808}, false},
		{`{"foo":-9223372036854775807.0}`, new(map[string]Int64), &map[string]Int64{"foo": -9223372036854775807}, false},
		{`{"foo":-9223372036854775806.0}`, new(map[string]Int64), &map[string]Int64{"foo": -9223372036854775806}, false},
		{`{"foo":9223372036854775805.0}`, new(map[string]Int64), &map[string]Int64{"foo": 9223372036854775805}, false},
		{`{"foo":9223372036854775806.0}`, new(map[string]Int64), &map[string]Int64{"foo": 9223372036854775806}, false},
		{`{"foo":9223372036854775807.0}`, new(map[string]Int64), &map[string]Int64{"foo": 9223372036854775807}, false},
		{`{"foo":9223372036854775808.0}`, new(map[string]Int64), nil, true},

		// range check for uint8.
		{`{"foo":-1}`, new(map[string]Uint8), nil, true},
		{`{"foo":0}`, new(map[string]Uint8), &map[string]Uint8{"foo": 0}, false},
		{`{"foo":255}`, new(map[string]Uint8), &map[string]Uint8{"foo": 255}, false},
		{`{"foo":256}`, new(map[string]Uint8), nil, true},

		// range check for uint16.
		{`{"foo":-1}`, new(map[string]Uint16), nil, true},
		{`{"foo":0}`, new(map[string]Uint16), &map[string]Uint16{"foo": 0}, false},
		{`{"foo":65535}`, new(map[string]Uint16), &map[string]Uint16{"foo": 65535}, false},
		{`{"foo":65536}`, new(map[string]Uint16), nil, true},

		// range check for uint32.
		{`{"foo":-1}`, new(map[string]Uint32), nil, true},
		{`{"foo":0}`, new(map[string]Uint32), &map[string]Uint32{"foo": 0}, false},
		{`{"foo":4294967295}`, new(map[string]Uint32), &map[string]Uint32{"foo": 4294967295}, false},
		{`{"foo":4294967296}`, new(map[string]Uint32), nil, true},
		{`{"foo":4294967295.0}`, new(map[string]Uint32), &map[string]Uint32{"foo": 4294967295}, false},
		{`{"foo":4294967296.0}`, new(map[string]Uint32), nil, true},

		// range check for uint64.
		{`{"foo":-1}`, new(map[string]Uint64), nil, true},
		{`{"foo":0}`, new(map[string]Uint64), &map[string]Uint64{"foo": 0}, false},
		{`{"foo":18446744073709551613}`, new(map[string]Uint64), &map[string]Uint64{"foo": 18446744073709551613}, false},
		{`{"foo":18446744073709551614}`, new(map[string]Uint64), &map[string]Uint64{"foo": 18446744073709551614}, false},
		{`{"foo":18446744073709551615}`, new(map[string]Uint64), &map[string]Uint64{"foo": 18446744073709551615}, false},
		{`{"foo":18446744073709551616}`, new(map[string]Uint64), nil, true},
		{`{"foo":18446744073709551613.0}`, new(map[string]Uint64), &map[string]Uint64{"foo": 18446744073709551613}, false},
		{`{"foo":18446744073709551614.0}`, new(map[string]Uint64), &map[string]Uint64{"foo": 18446744073709551614}, false},
		{`{"foo":18446744073709551615.0}`, new(map[string]Uint64), &map[string]Uint64{"foo": 18446744073709551615}, false},
		{`{"foo":18446744073709551616.0}`, new(map[string]Uint64), nil, true},
	}

	for _, tc := range testcases {
		err := json.Unmarshal([]byte(tc.in), tc.ptr)
		if tc.err {
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			continue
		}
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(tc.ptr, tc.out) {
			t.Errorf("%#v: got %#v, want %#v", tc.in, tc.ptr, tc.out)
		}
	}
}

package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalInt8_Generated(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{
			struct {
				Foo Int8 `json:"foo"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Int8 `json:"foo,omitempty"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Int8 `json:"foo"`
			}{0},
			`{"foo":0}`,
		},
		{
			struct {
				Foo Int8 `json:"foo,omitempty"`
			}{0},
			`{}`,
		},

		{
			struct {
				Foo *Int8 `json:"foo"`
			}{PtrInt8(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Int8 `json:"foo,omitempty"`
			}{PtrInt8(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Int8 `json:"foo"`
			}{PtrInt8(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Int8 `json:"foo,omitempty"`
			}{PtrInt8(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Int8 `json:"foo"`
			}{nil},
			`{"foo":null}`,
		},
		{
			struct {
				Foo *Int8 `json:"foo,omitempty"`
			}{nil},
			`{}`,
		},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.in)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if string(data) != tc.out {
			t.Errorf("%#v: got %#v, want %#v", tc.in, string(data), tc.out)
		}
	}
}

func TestUnmarshalInt8_Generated(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Int8), &map[string]Int8{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Int8), &map[string]*Int8{"foo": PtrInt8(123)}},
		{`{"foo":"123"}`, new(map[string]Int8), &map[string]Int8{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Int8), &map[string]*Int8{"foo": PtrInt8(123)}},
		{`{"foo":null}`, new(map[string]Int8), &map[string]Int8{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Int8), &map[string]*Int8{"foo": nil}},
		{`{"foo":""}`, new(map[string]Int8), &map[string]Int8{"foo": 0}},
		{`{"foo":""}`, new(map[string]*Int8), &map[string]*Int8{"foo": PtrInt8(0)}},
		{`{"foo":[123,"45",null]}`, new(map[string][]Int8), &map[string][]Int8{"foo": {123, 45, 0}}},
		{`{"foo":[123,"45",null]}`, new(map[string][]*Int8), &map[string][]*Int8{"foo": {PtrInt8(123), PtrInt8(45), nil}}},
	}

	for _, tc := range testcases {
		err := json.Unmarshal([]byte(tc.in), tc.ptr)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(tc.ptr, tc.out) {
			t.Errorf("%#v: got %#v, want %#v", tc.in, tc.ptr, tc.out)
		}
	}
}

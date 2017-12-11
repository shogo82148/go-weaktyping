package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalUint64_Generated(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{
			struct {
				Foo Uint64 `json:"foo"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Uint64 `json:"foo,omitempty"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Uint64 `json:"foo"`
			}{0},
			`{"foo":0}`,
		},
		{
			struct {
				Foo Uint64 `json:"foo,omitempty"`
			}{0},
			`{}`,
		},

		{
			struct {
				Foo *Uint64 `json:"foo"`
			}{PtrUint64(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Uint64 `json:"foo,omitempty"`
			}{PtrUint64(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Uint64 `json:"foo"`
			}{PtrUint64(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Uint64 `json:"foo,omitempty"`
			}{PtrUint64(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Uint64 `json:"foo"`
			}{nil},
			`{"foo":null}`,
		},
		{
			struct {
				Foo *Uint64 `json:"foo,omitempty"`
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

func TestUnmarshalUint64_Generated(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Uint64), &map[string]Uint64{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Uint64), &map[string]*Uint64{"foo": PtrUint64(123)}},
		{`{"foo":"123"}`, new(map[string]Uint64), &map[string]Uint64{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Uint64), &map[string]*Uint64{"foo": PtrUint64(123)}},
		{`{"foo":null}`, new(map[string]Uint64), &map[string]Uint64{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Uint64), &map[string]*Uint64{"foo": nil}},
		{`{"foo":[123,"45",null]}`, new(map[string][]Uint64), &map[string][]Uint64{"foo": {123, 45, 0}}},
		{`{"foo":[123,"45",null]}`, new(map[string][]*Uint64), &map[string][]*Uint64{"foo": {PtrUint64(123), PtrUint64(45), nil}}},
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

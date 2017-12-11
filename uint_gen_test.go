package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalUint_Generated(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{
			struct {
				Foo Uint `json:"foo"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Uint `json:"foo,omitempty"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Uint `json:"foo"`
			}{0},
			`{"foo":0}`,
		},
		{
			struct {
				Foo Uint `json:"foo,omitempty"`
			}{0},
			`{}`,
		},

		{
			struct {
				Foo *Uint `json:"foo"`
			}{PtrUint(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Uint `json:"foo,omitempty"`
			}{PtrUint(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Uint `json:"foo"`
			}{PtrUint(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Uint `json:"foo,omitempty"`
			}{PtrUint(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Uint `json:"foo"`
			}{nil},
			`{"foo":null}`,
		},
		{
			struct {
				Foo *Uint `json:"foo,omitempty"`
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

func TestUnmarshalUint_Generated(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Uint), &map[string]Uint{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Uint), &map[string]*Uint{"foo": PtrUint(123)}},
		{`{"foo":"123"}`, new(map[string]Uint), &map[string]Uint{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Uint), &map[string]*Uint{"foo": PtrUint(123)}},
		{`{"foo":null}`, new(map[string]Uint), &map[string]Uint{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Uint), &map[string]*Uint{"foo": nil}},
		{`{"foo":[123,"45",null]}`, new(map[string][]Uint), &map[string][]Uint{"foo": {123, 45, 0}}},
		{`{"foo":[123,"45",null]}`, new(map[string][]*Uint), &map[string][]*Uint{"foo": {PtrUint(123), PtrUint(45), nil}}},
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

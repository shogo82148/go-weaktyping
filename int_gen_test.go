package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalInt_Generated(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{
			struct {
				Foo Int `json:"foo"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Int `json:"foo,omitempty"`
			}{123},
			`{"foo":123}`,
		},
		{
			struct {
				Foo Int `json:"foo"`
			}{0},
			`{"foo":0}`,
		},
		{
			struct {
				Foo Int `json:"foo,omitempty"`
			}{0},
			`{}`,
		},

		{
			struct {
				Foo *Int `json:"foo"`
			}{PtrInt(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Int `json:"foo,omitempty"`
			}{PtrInt(123)},
			`{"foo":123}`,
		},
		{
			struct {
				Foo *Int `json:"foo"`
			}{PtrInt(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Int `json:"foo,omitempty"`
			}{PtrInt(0)},
			`{"foo":0}`,
		},
		{
			struct {
				Foo *Int `json:"foo"`
			}{nil},
			`{"foo":null}`,
		},
		{
			struct {
				Foo *Int `json:"foo,omitempty"`
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

func TestUnmarshalInt_Generated(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Int), &map[string]Int{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Int), &map[string]*Int{"foo": PtrInt(123)}},
		{`{"foo":"123"}`, new(map[string]Int), &map[string]Int{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Int), &map[string]*Int{"foo": PtrInt(123)}},
		{`{"foo":null}`, new(map[string]Int), &map[string]Int{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Int), &map[string]*Int{"foo": nil}},
		{`{"foo":""}`, new(map[string]Int), &map[string]Int{"foo": 0}},
		{`{"foo":""}`, new(map[string]*Int), &map[string]*Int{"foo": PtrInt(0)}},
		{`{"foo":[123,"45",null]}`, new(map[string][]Int), &map[string][]Int{"foo": {123, 45, 0}}},
		{`{"foo":[123,"45",null]}`, new(map[string][]*Int), &map[string][]*Int{"foo": {PtrInt(123), PtrInt(45), nil}}},
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

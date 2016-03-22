package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalInt16(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo Int16 `json:"foo"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Int16 `json:"foo,omitempty"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Int16 `json:"foo"`
		}{0}, `{"foo":0}`},
		{struct {
			Foo Int16 `json:"foo,omitempty"`
		}{0}, `{}`},

		{struct {
			Foo *Int16 `json:"foo"`
		}{PtrInt16(123)}, `{"foo":123}`},
		{struct {
			Foo *Int16 `json:"foo,omitempty"`
		}{PtrInt16(123)}, `{"foo":123}`},
		{struct {
			Foo *Int16 `json:"foo"`
		}{PtrInt16(0)}, `{"foo":0}`},
		{struct {
			Foo *Int16 `json:"foo,omitempty"`
		}{PtrInt16(0)}, `{"foo":0}`},
		{struct {
			Foo *Int16 `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *Int16 `json:"foo,omitempty"`
		}{nil}, `{}`},
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

func TestMarshalInt16(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Int16), &map[string]Int16{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Int16), &map[string]*Int16{"foo": PtrInt16(123)}},
		{`{"foo":"123"}`, new(map[string]Int16), &map[string]Int16{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Int16), &map[string]*Int16{"foo": PtrInt16(123)}},
		{`{"foo":null}`, new(map[string]Int16), &map[string]Int16{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Int16), &map[string]*Int16{"foo": nil}},
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
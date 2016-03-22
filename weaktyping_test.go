package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo Int `json:"foo"`
		}{Int{123}}, `{"foo":123}`},
		{struct {
			Foo Int `json:"foo,omitempty"`
		}{Int{123}}, `{"foo":123}`},
		{struct {
			Foo Int `json:"foo"`
		}{Int{}}, `{"foo":0}`},
		{struct {
			Foo Int `json:"foo,omitempty"`
		}{Int{}}, `{"foo":0}`},

		{struct {
			Foo *Int `json:"foo"`
		}{&Int{123}}, `{"foo":123}`},
		{struct {
			Foo *Int `json:"foo,omitempty"`
		}{&Int{123}}, `{"foo":123}`},
		{struct {
			Foo *Int `json:"foo"`
		}{&Int{}}, `{"foo":0}`},
		{struct {
			Foo *Int `json:"foo,omitempty"`
		}{&Int{}}, `{"foo":0}`},
		{struct {
			Foo *Int `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *Int `json:"foo,omitempty"`
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

func TestMarshal(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Int), &map[string]Int{"foo": Int{123}}},
		{`{"foo":123}`, new(map[string]*Int), &map[string]*Int{"foo": &Int{123}}},
		{`{"foo":"123"}`, new(map[string]Int), &map[string]Int{"foo": Int{123}}},
		{`{"foo":"123"}`, new(map[string]*Int), &map[string]*Int{"foo": &Int{123}}},
		{`{"foo":null}`, new(map[string]Int), &map[string]Int{"foo": Int{0}}},
		{`{"foo":null}`, new(map[string]*Int), &map[string]*Int{"foo": nil}},
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

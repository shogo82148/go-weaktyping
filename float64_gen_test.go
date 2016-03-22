package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalFloat64(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo Float64 `json:"foo"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Float64 `json:"foo,omitempty"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Float64 `json:"foo"`
		}{0}, `{"foo":0}`},
		{struct {
			Foo Float64 `json:"foo,omitempty"`
		}{0}, `{}`},

		{struct {
			Foo *Float64 `json:"foo"`
		}{PtrFloat64(123)}, `{"foo":123}`},
		{struct {
			Foo *Float64 `json:"foo,omitempty"`
		}{PtrFloat64(123)}, `{"foo":123}`},
		{struct {
			Foo *Float64 `json:"foo"`
		}{PtrFloat64(0)}, `{"foo":0}`},
		{struct {
			Foo *Float64 `json:"foo,omitempty"`
		}{PtrFloat64(0)}, `{"foo":0}`},
		{struct {
			Foo *Float64 `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *Float64 `json:"foo,omitempty"`
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

func TestMarshalFloat64(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Float64), &map[string]Float64{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Float64), &map[string]*Float64{"foo": PtrFloat64(123)}},
		{`{"foo":"123"}`, new(map[string]Float64), &map[string]Float64{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Float64), &map[string]*Float64{"foo": PtrFloat64(123)}},
		{`{"foo":null}`, new(map[string]Float64), &map[string]Float64{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Float64), &map[string]*Float64{"foo": nil}},
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
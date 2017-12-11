package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalFloat32_Generated(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo Float32 `json:"foo"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Float32 `json:"foo,omitempty"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Float32 `json:"foo"`
		}{0}, `{"foo":0}`},
		{struct {
			Foo Float32 `json:"foo,omitempty"`
		}{0}, `{}`},

		{struct {
			Foo *Float32 `json:"foo"`
		}{PtrFloat32(123)}, `{"foo":123}`},
		{struct {
			Foo *Float32 `json:"foo,omitempty"`
		}{PtrFloat32(123)}, `{"foo":123}`},
		{struct {
			Foo *Float32 `json:"foo"`
		}{PtrFloat32(0)}, `{"foo":0}`},
		{struct {
			Foo *Float32 `json:"foo,omitempty"`
		}{PtrFloat32(0)}, `{"foo":0}`},
		{struct {
			Foo *Float32 `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *Float32 `json:"foo,omitempty"`
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

func TestUnmarshalFloat32_Generated(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Float32), &map[string]Float32{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Float32), &map[string]*Float32{"foo": PtrFloat32(123)}},
		{`{"foo":"123"}`, new(map[string]Float32), &map[string]Float32{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Float32), &map[string]*Float32{"foo": PtrFloat32(123)}},
		{`{"foo":null}`, new(map[string]Float32), &map[string]Float32{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Float32), &map[string]*Float32{"foo": nil}},
		{`{"foo":[123,"45",null]}`, new(map[string][]Float32), &map[string][]Float32{"foo": {123, 45, 0}}},
		{`{"foo":[123,"45",null]}`, new(map[string][]*Float32), &map[string][]*Float32{"foo": {PtrFloat32(123), PtrFloat32(45), nil}}},
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

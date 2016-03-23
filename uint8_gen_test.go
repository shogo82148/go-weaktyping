package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalUint8(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo Uint8 `json:"foo"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Uint8 `json:"foo,omitempty"`
		}{123}, `{"foo":123}`},
		{struct {
			Foo Uint8 `json:"foo"`
		}{0}, `{"foo":0}`},
		{struct {
			Foo Uint8 `json:"foo,omitempty"`
		}{0}, `{}`},

		{struct {
			Foo *Uint8 `json:"foo"`
		}{PtrUint8(123)}, `{"foo":123}`},
		{struct {
			Foo *Uint8 `json:"foo,omitempty"`
		}{PtrUint8(123)}, `{"foo":123}`},
		{struct {
			Foo *Uint8 `json:"foo"`
		}{PtrUint8(0)}, `{"foo":0}`},
		{struct {
			Foo *Uint8 `json:"foo,omitempty"`
		}{PtrUint8(0)}, `{"foo":0}`},
		{struct {
			Foo *Uint8 `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *Uint8 `json:"foo,omitempty"`
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

func TestMarshalUint8(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]Uint8), &map[string]Uint8{"foo": 123}},
		{`{"foo":123}`, new(map[string]*Uint8), &map[string]*Uint8{"foo": PtrUint8(123)}},
		{`{"foo":"123"}`, new(map[string]Uint8), &map[string]Uint8{"foo": 123}},
		{`{"foo":"123"}`, new(map[string]*Uint8), &map[string]*Uint8{"foo": PtrUint8(123)}},
		{`{"foo":null}`, new(map[string]Uint8), &map[string]Uint8{"foo": 0}},
		{`{"foo":null}`, new(map[string]*Uint8), &map[string]*Uint8{"foo": nil}},
		{`{"foo":[123,"45",null]}`, new(map[string][]Uint8), &map[string][]Uint8{"foo": {123, 45, 0}}},
		{`{"foo":[123,"45",null]}`, new(map[string][]*Uint8), &map[string][]*Uint8{"foo": {PtrUint8(123), PtrUint8(45), nil}}},
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

package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalString(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo String `json:"foo"`
		}{"123"}, `{"foo":"123"}`},
		{struct {
			Foo String `json:"foo,omitempty"`
		}{"123"}, `{"foo":"123"}`},
		{struct {
			Foo String `json:"foo"`
		}{""}, `{"foo":""}`},
		{struct {
			Foo String `json:"foo,omitempty"`
		}{""}, `{}`},

		{struct {
			Foo *String `json:"foo"`
		}{PtrString("123")}, `{"foo":"123"}`},
		{struct {
			Foo *String `json:"foo,omitempty"`
		}{PtrString("123")}, `{"foo":"123"}`},
		{struct {
			Foo *String `json:"foo"`
		}{PtrString("")}, `{"foo":""}`},
		{struct {
			Foo *String `json:"foo,omitempty"`
		}{PtrString("")}, `{"foo":""}`},
		{struct {
			Foo *String `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *String `json:"foo,omitempty"`
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

func TestMarshalString(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":123}`, new(map[string]String), &map[string]String{"foo": "123"}},
		{`{"foo":123}`, new(map[string]*String), &map[string]*String{"foo": PtrString("123")}},
		{`{"foo":3.14}`, new(map[string]String), &map[string]String{"foo": "3.14"}},
		{`{"foo":3.14}`, new(map[string]*String), &map[string]*String{"foo": PtrString("3.14")}},
		{`{"foo":"123"}`, new(map[string]String), &map[string]String{"foo": "123"}},
		{`{"foo":"123"}`, new(map[string]*String), &map[string]*String{"foo": PtrString("123")}},
		{`{"foo":"\u3042"}`, new(map[string]String), &map[string]String{"foo": "\u3042"}},
		{`{"foo":"\u3042"}`, new(map[string]*String), &map[string]*String{"foo": PtrString("\u3042")}},
		{`{"foo":null}`, new(map[string]String), &map[string]String{"foo": "null"}},
		{`{"foo":null}`, new(map[string]*String), &map[string]*String{"foo": nil}},
		{`{"foo":true}`, new(map[string]String), &map[string]String{"foo": "true"}},
		{`{"foo":true}`, new(map[string]*String), &map[string]*String{"foo": PtrString("true")}},
		{`{"foo":false}`, new(map[string]String), &map[string]String{"foo": "false"}},
		{`{"foo":false}`, new(map[string]*String), &map[string]*String{"foo": PtrString("false")}},
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

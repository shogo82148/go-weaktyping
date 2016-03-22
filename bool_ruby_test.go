package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalRubyBool(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo RubyBool `json:"foo"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo RubyBool `json:"foo,omitempty"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo RubyBool `json:"foo"`
		}{false}, `{"foo":false}`},
		{struct {
			Foo RubyBool `json:"foo,omitempty"`
		}{false}, `{}`},

		{struct {
			Foo *RubyBool `json:"foo"`
		}{PtrRubyBool(true)}, `{"foo":true}`},
		{struct {
			Foo *RubyBool `json:"foo,omitempty"`
		}{PtrRubyBool(true)}, `{"foo":true}`},
		{struct {
			Foo *RubyBool `json:"foo"`
		}{PtrRubyBool(false)}, `{"foo":false}`},
		{struct {
			Foo *RubyBool `json:"foo,omitempty"`
		}{PtrRubyBool(false)}, `{"foo":false}`},
		{struct {
			Foo *RubyBool `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *RubyBool `json:"foo,omitempty"`
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

func TestMarshalRubyBool(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":true}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": true}},
		{`{"foo":true}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(true)}},
		{`{"foo":false}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": false}},
		{`{"foo":false}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(false)}},
		{`{"foo":1}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": true}},
		{`{"foo":1}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(true)}},
		{`{"foo":0}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": true}},
		{`{"foo":0}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(true)}},
		{`{"foo":0.0}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": true}},
		{`{"foo":0.0}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(true)}},
		{`{"foo":"0"}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": true}},
		{`{"foo":"0"}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(true)}},
		{`{"foo":""}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": true}},
		{`{"foo":""}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(true)}},
		{`{"foo":[]}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": true}},
		{`{"foo":[]}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": PtrRubyBool(true)}},
		{`{"foo":null}`, new(map[string]RubyBool), &map[string]RubyBool{"foo": false}},
		{`{"foo":null}`, new(map[string]*RubyBool), &map[string]*RubyBool{"foo": nil}},
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

package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalPerlBool(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo PerlBool `json:"foo"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo PerlBool `json:"foo,omitempty"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo PerlBool `json:"foo"`
		}{false}, `{"foo":false}`},
		{struct {
			Foo PerlBool `json:"foo,omitempty"`
		}{false}, `{}`},

		{struct {
			Foo *PerlBool `json:"foo"`
		}{PtrPerlBool(true)}, `{"foo":true}`},
		{struct {
			Foo *PerlBool `json:"foo,omitempty"`
		}{PtrPerlBool(true)}, `{"foo":true}`},
		{struct {
			Foo *PerlBool `json:"foo"`
		}{PtrPerlBool(false)}, `{"foo":false}`},
		{struct {
			Foo *PerlBool `json:"foo,omitempty"`
		}{PtrPerlBool(false)}, `{"foo":false}`},
		{struct {
			Foo *PerlBool `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *PerlBool `json:"foo,omitempty"`
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

func TestMarshalPerlBool(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":true}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": true}},
		{`{"foo":true}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(true)}},
		{`{"foo":false}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": false}},
		{`{"foo":false}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(false)}},
		{`{"foo":1}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": true}},
		{`{"foo":1}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(true)}},
		{`{"foo":0}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": false}},
		{`{"foo":0}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(false)}},
		{`{"foo":0.0}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": false}},
		{`{"foo":0.0}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(false)}},
		{`{"foo":"0"}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": false}},
		{`{"foo":"0"}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(false)}},
		{`{"foo":""}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": false}},
		{`{"foo":""}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(false)}},
		{`{"foo":[]}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": true}},
		{`{"foo":[]}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": PtrPerlBool(true)}},
		{`{"foo":null}`, new(map[string]PerlBool), &map[string]PerlBool{"foo": false}},
		{`{"foo":null}`, new(map[string]*PerlBool), &map[string]*PerlBool{"foo": nil}},
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

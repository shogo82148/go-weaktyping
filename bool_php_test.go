package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalPHPBool(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo PHPBool `json:"foo"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo PHPBool `json:"foo,omitempty"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo PHPBool `json:"foo"`
		}{false}, `{"foo":false}`},
		{struct {
			Foo PHPBool `json:"foo,omitempty"`
		}{false}, `{}`},

		{struct {
			Foo *PHPBool `json:"foo"`
		}{PtrPHPBool(true)}, `{"foo":true}`},
		{struct {
			Foo *PHPBool `json:"foo,omitempty"`
		}{PtrPHPBool(true)}, `{"foo":true}`},
		{struct {
			Foo *PHPBool `json:"foo"`
		}{PtrPHPBool(false)}, `{"foo":false}`},
		{struct {
			Foo *PHPBool `json:"foo,omitempty"`
		}{PtrPHPBool(false)}, `{"foo":false}`},
		{struct {
			Foo *PHPBool `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *PHPBool `json:"foo,omitempty"`
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

func TestMarshalPHPBool(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":true}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": true}},
		{`{"foo":true}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(true)}},
		{`{"foo":false}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": false}},
		{`{"foo":false}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(false)}},
		{`{"foo":1}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": true}},
		{`{"foo":1}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(true)}},
		{`{"foo":0}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": false}},
		{`{"foo":0}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(false)}},
		{`{"foo":0.0}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": false}},
		{`{"foo":0.0}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(false)}},
		{`{"foo":"0"}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": false}},
		{`{"foo":"0"}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(false)}},
		{`{"foo":""}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": false}},
		{`{"foo":""}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(false)}},
		{`{"foo":[]}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": false}},
		{`{"foo":[]}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": PtrPHPBool(false)}},
		{`{"foo":null}`, new(map[string]PHPBool), &map[string]PHPBool{"foo": false}},
		{`{"foo":null}`, new(map[string]*PHPBool), &map[string]*PHPBool{"foo": nil}},
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

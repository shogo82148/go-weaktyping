package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalPythonBool(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo PythonBool `json:"foo"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo PythonBool `json:"foo,omitempty"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo PythonBool `json:"foo"`
		}{false}, `{"foo":false}`},
		{struct {
			Foo PythonBool `json:"foo,omitempty"`
		}{false}, `{}`},

		{struct {
			Foo *PythonBool `json:"foo"`
		}{PtrPythonBool(true)}, `{"foo":true}`},
		{struct {
			Foo *PythonBool `json:"foo,omitempty"`
		}{PtrPythonBool(true)}, `{"foo":true}`},
		{struct {
			Foo *PythonBool `json:"foo"`
		}{PtrPythonBool(false)}, `{"foo":false}`},
		{struct {
			Foo *PythonBool `json:"foo,omitempty"`
		}{PtrPythonBool(false)}, `{"foo":false}`},
		{struct {
			Foo *PythonBool `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *PythonBool `json:"foo,omitempty"`
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

func TestMarshalPythonBool(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":true}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": true}},
		{`{"foo":true}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(true)}},
		{`{"foo":false}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": false}},
		{`{"foo":false}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(false)}},
		{`{"foo":1}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": true}},
		{`{"foo":1}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(true)}},
		{`{"foo":0}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": false}},
		{`{"foo":0}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(false)}},
		{`{"foo":0.0}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": false}},
		{`{"foo":0.0}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(false)}},
		{`{"foo":"0"}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": true}},
		{`{"foo":"0"}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(true)}},
		{`{"foo":""}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": false}},
		{`{"foo":""}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(false)}},
		{`{"foo":[]}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": false}},
		{`{"foo":[]}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(false)}},
		{`{"foo":{}}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": false}},
		{`{"foo":{}}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": PtrPythonBool(false)}},
		{`{"foo":null}`, new(map[string]PythonBool), &map[string]PythonBool{"foo": false}},
		{`{"foo":null}`, new(map[string]*PythonBool), &map[string]*PythonBool{"foo": nil}},
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

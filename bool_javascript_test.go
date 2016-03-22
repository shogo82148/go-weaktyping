package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalJavaScriptBool(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo JavaScriptBool `json:"foo"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo JavaScriptBool `json:"foo,omitempty"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo JavaScriptBool `json:"foo"`
		}{false}, `{"foo":false}`},
		{struct {
			Foo JavaScriptBool `json:"foo,omitempty"`
		}{false}, `{}`},

		{struct {
			Foo *JavaScriptBool `json:"foo"`
		}{PtrJavaScriptBool(true)}, `{"foo":true}`},
		{struct {
			Foo *JavaScriptBool `json:"foo,omitempty"`
		}{PtrJavaScriptBool(true)}, `{"foo":true}`},
		{struct {
			Foo *JavaScriptBool `json:"foo"`
		}{PtrJavaScriptBool(false)}, `{"foo":false}`},
		{struct {
			Foo *JavaScriptBool `json:"foo,omitempty"`
		}{PtrJavaScriptBool(false)}, `{"foo":false}`},
		{struct {
			Foo *JavaScriptBool `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *JavaScriptBool `json:"foo,omitempty"`
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

func TestMarshalJavaScriptBool(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":true}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": true}},
		{`{"foo":true}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(true)}},
		{`{"foo":false}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": false}},
		{`{"foo":false}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(false)}},
		{`{"foo":1}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": true}},
		{`{"foo":1}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(true)}},
		{`{"foo":0}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": false}},
		{`{"foo":0}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(false)}},
		{`{"foo":0.0}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": false}},
		{`{"foo":0.0}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(false)}},
		{`{"foo":"0"}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": true}},
		{`{"foo":"0"}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(true)}},
		{`{"foo":""}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": false}},
		{`{"foo":""}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(false)}},
		{`{"foo":[]}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": true}},
		{`{"foo":[]}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": PtrJavaScriptBool(true)}},
		{`{"foo":null}`, new(map[string]JavaScriptBool), &map[string]JavaScriptBool{"foo": false}},
		{`{"foo":null}`, new(map[string]*JavaScriptBool), &map[string]*JavaScriptBool{"foo": nil}},
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

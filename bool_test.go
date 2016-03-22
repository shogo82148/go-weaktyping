package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalBool(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo Bool `json:"foo"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo Bool `json:"foo,omitempty"`
		}{true}, `{"foo":true}`},
		{struct {
			Foo Bool `json:"foo"`
		}{false}, `{"foo":false}`},
		{struct {
			Foo Bool `json:"foo,omitempty"`
		}{false}, `{}`},

		{struct {
			Foo *Bool `json:"foo"`
		}{PtrBool(true)}, `{"foo":true}`},
		{struct {
			Foo *Bool `json:"foo,omitempty"`
		}{PtrBool(true)}, `{"foo":true}`},
		{struct {
			Foo *Bool `json:"foo"`
		}{PtrBool(false)}, `{"foo":false}`},
		{struct {
			Foo *Bool `json:"foo,omitempty"`
		}{PtrBool(false)}, `{"foo":false}`},
		{struct {
			Foo *Bool `json:"foo"`
		}{nil}, `{"foo":null}`},
		{struct {
			Foo *Bool `json:"foo,omitempty"`
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

func TestMarshalBool(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{`{"foo":true}`, new(map[string]Bool), &map[string]Bool{"foo": true}},
		{`{"foo":true}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(true)}},
		{`{"foo":false}`, new(map[string]Bool), &map[string]Bool{"foo": false}},
		{`{"foo":false}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(false)}},
		{`{"foo":1}`, new(map[string]Bool), &map[string]Bool{"foo": true}},
		{`{"foo":1}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(true)}},
		{`{"foo":0}`, new(map[string]Bool), &map[string]Bool{"foo": false}},
		{`{"foo":0}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(false)}},
		{`{"foo":0.0}`, new(map[string]Bool), &map[string]Bool{"foo": false}},
		{`{"foo":0.0}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(false)}},
		{`{"foo":"0"}`, new(map[string]Bool), &map[string]Bool{"foo": true}},
		{`{"foo":"0"}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(true)}},
		{`{"foo":""}`, new(map[string]Bool), &map[string]Bool{"foo": false}},
		{`{"foo":""}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(false)}},
		{`{"foo":[]}`, new(map[string]Bool), &map[string]Bool{"foo": true}},
		{`{"foo":[]}`, new(map[string]*Bool), &map[string]*Bool{"foo": PtrBool(true)}},
		{`{"foo":null}`, new(map[string]Bool), &map[string]Bool{"foo": false}},
		{`{"foo":null}`, new(map[string]*Bool), &map[string]*Bool{"foo": nil}},
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

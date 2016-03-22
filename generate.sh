#!/bin/sh

set -ue

NEW_TYPE="$1"
ORG_TYPE="$2"
PARSE="$3"

cat <<EOF > "${ORG_TYPE}_gen.go"
package weaktyping

import "strconv"

type $NEW_TYPE $ORG_TYPE

func Ptr$NEW_TYPE(v $NEW_TYPE) *$NEW_TYPE {
	return &v
}

func (v *$NEW_TYPE) UnmarshalJSON(data []byte) error {
	return v.UnmarshalText(unquoteBytesIfQuoted(data))
}

func (v *$NEW_TYPE) UnmarshalText(data []byte) error {
	s := string(data)
	if s == "null" {
		*v = 0
		return nil
	}
	if w, err := $PARSE; err != nil {
		return err
	} else {
		*v = $NEW_TYPE(w)
	}
	return nil
}
EOF

go fmt "${ORG_TYPE}_gen.go"

cat <<EOF > "${ORG_TYPE}_gen_test.go"
package weaktyping

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshal$NEW_TYPE(t *testing.T) {
	testcases := []struct {
		in  interface{}
		out string
	}{
		{struct {
			Foo $NEW_TYPE \`json:"foo"\`
		}{123}, \`{"foo":123}\`},
		{struct {
			Foo $NEW_TYPE \`json:"foo,omitempty"\`
		}{123}, \`{"foo":123}\`},
		{struct {
			Foo $NEW_TYPE \`json:"foo"\`
		}{0}, \`{"foo":0}\`},
		{struct {
			Foo $NEW_TYPE \`json:"foo,omitempty"\`
		}{0}, \`{}\`},

		{struct {
			Foo *$NEW_TYPE \`json:"foo"\`
		}{Ptr$NEW_TYPE(123)}, \`{"foo":123}\`},
		{struct {
			Foo *$NEW_TYPE \`json:"foo,omitempty"\`
		}{Ptr$NEW_TYPE(123)}, \`{"foo":123}\`},
		{struct {
			Foo *$NEW_TYPE \`json:"foo"\`
		}{Ptr$NEW_TYPE(0)}, \`{"foo":0}\`},
		{struct {
			Foo *$NEW_TYPE \`json:"foo,omitempty"\`
		}{Ptr$NEW_TYPE(0)}, \`{"foo":0}\`},
		{struct {
			Foo *$NEW_TYPE \`json:"foo"\`
		}{nil}, \`{"foo":null}\`},
		{struct {
			Foo *$NEW_TYPE \`json:"foo,omitempty"\`
		}{nil}, \`{}\`},
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

func TestMarshal$NEW_TYPE(t *testing.T) {
	testcases := []struct {
		in  string
		ptr interface{}
		out interface{}
	}{
		{\`{"foo":123}\`, new(map[string]$NEW_TYPE), &map[string]$NEW_TYPE{"foo": 123}},
		{\`{"foo":123}\`, new(map[string]*$NEW_TYPE), &map[string]*$NEW_TYPE{"foo": Ptr$NEW_TYPE(123)}},
		{\`{"foo":"123"}\`, new(map[string]$NEW_TYPE), &map[string]$NEW_TYPE{"foo": 123}},
		{\`{"foo":"123"}\`, new(map[string]*$NEW_TYPE), &map[string]*$NEW_TYPE{"foo": Ptr$NEW_TYPE(123)}},
		{\`{"foo":null}\`, new(map[string]$NEW_TYPE), &map[string]$NEW_TYPE{"foo": 0}},
		{\`{"foo":null}\`, new(map[string]*$NEW_TYPE), &map[string]*$NEW_TYPE{"foo": nil}},
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
EOF

go fmt "${ORG_TYPE}_gen_test.go"

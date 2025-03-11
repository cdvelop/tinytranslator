package tinytranslator

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	tests := []struct {
		tag      reflect.StructTag
		expected map[string]string
	}{
		{
			tag:      `json:"name" xml:"name"`,
			expected: map[string]string{"json": "name", "xml": "name"},
		},
		{
			tag:      `json:"age" xml:"age"`,
			expected: map[string]string{"json": "age", "xml": "age"},
		},
		{
			tag:      `db:"id" json:"id"`,
			expected: map[string]string{"db": "id", "json": "id"},
		},
		{
			tag:      `json:"-" xml:"-"`,
			expected: map[string]string{"json": "-", "xml": "-"},
		},
		{
			tag:      `json:"name,omitempty" xml:"name,omitempty"`,
			expected: map[string]string{"json": "name,omitempty", "xml": "name,omitempty"},
		},
	}

	for _, test := range tests {
		result := parseTag(test.tag)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("parseTag(%q) = %v; want %v", test.tag, result, test.expected)
		}
	}
}

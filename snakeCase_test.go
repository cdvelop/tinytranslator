package lang_test

import (
	"testing"

	"github.com/cesarsoliscaro/lang"
)

func TestSnakeCase(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"NameFirst", "name_first"},
		{"nameFirst", "name_first"},
		{"NameFirstSecond", "name_first_second"},
		{"_other_name", "_other_name"},
		{"123Test", "123_test"},
		{"", ""},
		{"ALLCAPS", "allcaps"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {

			got := lang.SnakeCase(tc.input)
			if got != tc.want {
				t.Fatalf("snakeCase(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

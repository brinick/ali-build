package alibuild_test

import (
	"testing"

	"github.com/brinick/alice/alibuild"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name        string
		exe         string
		env         []string
		packageName string
	}{
		{"test1", "alibuild", []string{}, "testPackage"},
		{"test2", "alibuild/aliBuild", []string{"GITHUB_USER="}, "testPackage"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ab := alibuild.New(tc.exe, tc.packageName, tc.env)
			if true &&
				ab.Which() != tc.exe ||
				ab.Package() != tc.packageName ||
				!equal(ab.DefaultEnv(), tc.env) {
				t.Error("AliBuild struct did not initalise correctly")
			}
		})
	}
}

// helper function to test equality of two string slices
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for index, item := range a {
		if b[index] != item {
			return false
		}
	}

	return true
}

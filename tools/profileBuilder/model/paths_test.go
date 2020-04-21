package model

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestDeconstructPath(t *testing.T) {
	type testcase struct {
		name string
		path string
		pi   PathInfo
	}
	testcases := []testcase{
		{
			name: "resource",
			path: filepath.Join("work", "src", "github.com", "samjegal", "fincloud-sdk-for-go", "services", "network"),
			pi: PathInfo{
				IsArm:    true,
				Provider: "network",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := DeconstructPath(tc.path)
			if err != nil {
				t.Fatalf("failed to deconstruct path: %v", err)
			}
			if !reflect.DeepEqual(p, tc.pi) {
				t.Fatalf("expected %+v, got %+v", tc.pi, p)
			}
		})
	}
}

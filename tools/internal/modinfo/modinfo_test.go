package modinfo

import (
	"regexp"
	"testing"
)

func Test_ScenarioA(t *testing.T) {
	mod, err := GetModuleInfo("../../testdata/scenarioa/foo", "../../testdata/scenarioa/foo/stage")
	if err != nil {
		t.Fatalf("failed to get module info: %v", err)
	}
	if mod.BreakingChanges() {
		t.Fatal("no breaking changes in scenario A")
	}
	if !mod.NewExports() {
		t.Fatal("expected new exports in scenario A")
	}
	if mod.VersionSuffix() {
		t.Fatalf("unexpected version suffix in scenario A")
	}
	regex := regexp.MustCompile(`testdata/scenarioa/foo$`)
	if !regex.MatchString(mod.DestDir()) {
		t.Fatalf("bad destination dir: %s", mod.DestDir())
	}
}

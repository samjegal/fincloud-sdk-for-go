package modinfo

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/samjegal/fincloud-sdk-for-go/tools/apidiff/delta"
	"github.com/samjegal/fincloud-sdk-for-go/tools/apidiff/exports"
	"github.com/samjegal/fincloud-sdk-for-go/tools/apidiff/report"
	"github.com/samjegal/fincloud-sdk-for-go/tools/internal/dirs"
)

var (
	verSuffixRegex = regexp.MustCompile(`v\d+$`)
)

func HasVersionSuffix(path string) bool {
	return verSuffixRegex.MatchString(path)
}

func FindVersionSuffix(path string) string {
	return verSuffixRegex.FindString(path)
}

func GetModuleSubdirs(path string) ([]string, error) {
	subdirs, err := dirs.GetSubdirs(path)
	if err != nil {
		return nil, err
	}
	modDirs := []string{}
	for _, subdir := range subdirs {
		matched := HasVersionSuffix(subdir)
		if matched {
			modDirs = append(modDirs, subdir)
		}
	}
	sortModuleTagsBySemver(modDirs)
	return modDirs, nil
}

func IncrementModuleVersion(ver string) string {
	if ver == "" {
		return "v2"
	}
	v, err := strconv.ParseInt(ver[1:], 10, 32)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("v%d", v+1)
}

func sortModuleTagsBySemver(modDirs []string) {
	sort.SliceStable(modDirs, func(i, j int) bool {
		lv, err := semver.NewVersion(modDirs[i])
		if err != nil {
			panic(err)
		}
		rv, err := semver.NewVersion(modDirs[j])
		if err != nil {
			panic(err)
		}
		return lv.LessThan(rv)
	})
}

func CreateModuleNameFromPath(pkgDir string) (string, error) {
	repoRoot := filepath.Join("github.com", "samjegal", "fincloud-sdk-for-go")
	i := strings.Index(pkgDir, repoRoot)
	if i < 0 {
		return "", fmt.Errorf("didn't find '%s' in '%s'", repoRoot, pkgDir)
	}
	return strings.Replace(pkgDir[i:], "\\", "/", -1), nil
}

type Provider interface {
	DestDir() string
	NewExports() bool
	BreakingChanges() bool
	VersionSuffix() bool
	NewModule() bool
	GenerateReport() report.Package
}

type module struct {
	lhs  exports.Content
	rhs  exports.Content
	dest string
}

func GetModuleInfo(baseline, staged string) (Provider, error) {
	// TODO: verify staged is a child of baseline
	lhs, err := exports.Get(baseline)
	if err != nil {
		// if baseline has no content then this is a v1 package
		if ei, ok := err.(exports.LoadPackageErrorInfo); !ok || ei.PkgCount() != 0 {
			return nil, fmt.Errorf("failed to get exports for package '%s': %s", baseline, err)
		}
	}
	rhs, err := exports.Get(staged)
	if err != nil {
		return nil, fmt.Errorf("failed to get exports for package '%s': %s", staged, err)
	}
	mod := module{
		lhs:  lhs,
		rhs:  rhs,
		dest: baseline,
	}
	// calculate the destination directory
	// if there are breaking changes calculate the new directory
	if mod.BreakingChanges() {
		dest := filepath.Dir(staged)
		v := 2
		if verSuffixRegex.MatchString(baseline) {
			// baseline has a version, get the number and increment it
			s := string(baseline[len(baseline)-1])
			v, err = strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("failed to convert '%s' to int: %v", s, err)
			}
			v++
		}
		mod.dest = filepath.Join(dest, fmt.Sprintf("v%d", v))
	}
	return mod, nil
}

func (m module) DestDir() string {
	return m.dest
}

func (m module) NewExports() bool {
	if m.lhs.IsEmpty() {
		return true
	}
	adds := delta.GetExports(m.lhs, m.rhs)
	return !adds.IsEmpty()
}

func (m module) BreakingChanges() bool {
	if m.lhs.IsEmpty() {
		return false
	}
	// check for changed content
	if len(delta.GetConstTypeChanges(m.lhs, m.rhs)) > 0 ||
		len(delta.GetFuncSigChanges(m.lhs, m.rhs)) > 0 ||
		len(delta.GetInterfaceMethodSigChanges(m.lhs, m.rhs)) > 0 ||
		len(delta.GetStructFieldChanges(m.lhs, m.rhs)) > 0 {
		return true
	}
	// check for removed content
	if removed := delta.GetExports(m.rhs, m.lhs); !removed.IsEmpty() {
		return true
	}
	return false
}

func (m module) VersionSuffix() bool {
	return verSuffixRegex.MatchString(m.dest)
}

func (m module) NewModule() bool {
	return m.lhs.IsEmpty()
}

func (m module) GenerateReport() report.Package {
	return report.Generate(m.lhs, m.rhs, false, false)
}

func IsValidModuleVersion(v string) bool {
	r := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	return r.MatchString(v)
}

package model

import (
	"fmt"
	"path/filepath"
	"regexp"
)

const pkgPathRegex = `services[/\\](?P<prv>[\w\-\._/\\]+)`

var (
	validPackagePath   = regexp.MustCompile(pkgPathRegex)
	validPkgPathModVer = regexp.MustCompile(pkgPathRegex + `[/\\](?P<mod>v\d+)[/\\]?(?P<api>\w+api)?`)
	apiPkgRegex        = regexp.MustCompile(`[/\\]\w+api$`)
)

type PathInfo struct {
	IsArm    bool
	Provider string
	Version  string
	Group    string
	ModVer   string
	APIPkg   string
}

func DeconstructPath(path string) (PathInfo, error) {
	path = filepath.Clean(path)
	matches := validPkgPathModVer.FindAllStringSubmatch(path, -1)
	if matches == nil {
		matches = validPackagePath.FindAllStringSubmatch(path, -1)
	}
	if matches == nil {
		return PathInfo{}, fmt.Errorf("path '%s' does not resemble a known package path", path)
	}

	prv := matches[0][1]
	mod := ""
	api := ""

	if len(matches[0]) == 4 {
		mod = matches[0][2]
		api = matches[0][3]
	}

	arm := true
	return PathInfo{
		Provider: prv,
		IsArm:    arm,
		ModVer:   mod,
		APIPkg:   api,
	}, nil
}

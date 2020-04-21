package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const previewSubdir = string(os.PathSeparator) + "preview" + string(os.PathSeparator)

var previewVer = regexp.MustCompile(`(?:v?\d{4}-\d{2}-\d{2}|v?\d+[\.\d+\.\d\-]*)(?:-preview|-beta)`)

func acceptAllPredicate(name string) bool {
	return true
}

func includePreviewPredicate(name string) bool {
	if strings.Contains(name, previewSubdir) {
		return false
	}
	return !previewVer.MatchString(name)
}

type operationGroup struct {
	provider string
	arm      bool
	group    string
	api      string
}

type operInfo struct {
	version string
	rawpath string
	modver  string
}

type latestTracker map[operationGroup]operInfo

func (tracker latestTracker) Upsert(path string, pi PathInfo) (operInfo, operationGroup, int) {
	group := operationGroup{
		provider: pi.Provider,
		arm:      pi.IsArm,
		group:    pi.Group,
		api:      pi.APIPkg,
	}

	prev, ok := tracker[group]
	if !ok {
		tracker[group] = operInfo{pi.Version, path, pi.ModVer}
		return prev, group, 1
	}

	if le, _ := versionLE(prev.version, pi.Version); le {
		tracker[group] = operInfo{pi.Version, path, pi.ModVer}
		return prev, group, 0
	}
	return prev, group, -1
}

func (tracker latestTracker) ToListDefinition(skipFileCheck bool) (ListDefinition, error) {
	listDef := ListDefinition{
		Include: []string{},
	}
	for _, entry := range tracker {
		if skipFileCheck {
			listDef.Include = append(listDef.Include, entry.rawpath)
			continue
		}
		absolute, err := filepath.Abs(entry.rawpath)
		if err != nil {
			return listDef, err
		}
		entries, err := ioutil.ReadDir(absolute)
		if err != nil {
			return listDef, err
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				listDef.Include = append(listDef.Include, absolute)
				break
			}
		}
	}
	sort.Strings(listDef.Include)
	return listDef, nil
}

func GetLatestPackages(rootDir string, includePreview bool, verboseLog *log.Logger) (ListDefinition, error) {
	predicate := includePreviewPredicate
	if includePreview {
		predicate = acceptAllPredicate
	}

	tracker := latestTracker{}

	filepath.Walk(rootDir, func(currentPath string, info os.FileInfo, openErr error) error {
		pi, err := DeconstructPath(currentPath)
		if err != nil || !info.IsDir() {
			return nil
		} else if !predicate(currentPath) {
			verboseLog.Printf("%q rejected by Predicate", currentPath)
			return nil
		}

		prev, group, result := tracker.Upsert(currentPath, pi)
		switch result {
		case 1:
			verboseLog.Printf("New group found %v using version %q modver %q", group, pi.Version, pi.ModVer)
		case 0:
			verboseLog.Printf("Updating group %v from version %q to %q modver %q", group, prev.version, pi.Version, pi.ModVer)
		case -1:
			verboseLog.Printf("Evaluated group %v version %q decided to stay with %q", group, pi.Version, prev.version)
		}

		return nil
	})

	return tracker.ToListDefinition(false)
}

var versionLE = func() func(string, string) (bool, error) {
	type strategyTuple struct {
		match   *regexp.Regexp
		handler func([]string, []string) (bool, error)
	}
	wellKnownStrategies := []strategyTuple{
		{
			match: regexp.MustCompile(`^(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})(?:[\.\-](?P<tag>.+))?$`),
			handler: func(leftMatch, rightMatch []string) (bool, error) {
				var err error
				for i := 1; i <= 3; i++ { // Start with index 1 because the element 0 is the entire match, not a group. End at 3 because there are three numeric groups.
					if leftMatch[i] == rightMatch[i] {
						continue
					}

					var leftNum, rightNum int
					leftNum, err = strconv.Atoi(leftMatch[i])
					if err != nil {
						return false, err
					}

					rightNum, err = strconv.Atoi(rightMatch[i])
					if err != nil {
						return false, err
					}

					if leftNum < rightNum {
						return true, nil
					}
					return false, nil
				}

				if leftTag, rightTag := leftMatch[4], rightMatch[4]; leftTag == "" && rightTag != "" { // match[4] is the tag portion of a date based API Version label
					return false, nil
				} else if leftTag != "" && rightTag != "" {
					return leftTag <= rightTag, nil
				}
				return true, nil
			},
		},
		{
			match: regexp.MustCompile(`(?P<major>\d+)(?:\.(?P<minor>\d+)(?:\.(?P<patch>\d+))?-?(?P<tag>.*))?`),
			handler: func(leftMatch, rightMatch []string) (bool, error) {
				for i := 1; i <= 3; i++ {
					if len(leftMatch[i]) == 0 || len(rightMatch[i]) == 0 {
						return leftMatch[i] <= rightMatch[i], nil
					}
					numLeft, err := strconv.Atoi(leftMatch[i])
					if err != nil {
						return false, err
					}
					numRight, err := strconv.Atoi(rightMatch[i])
					if err != nil {
						return false, err
					}

					if numLeft < numRight {
						return true, nil
					}

					if numLeft > numRight {
						return false, nil
					}
				}

				return leftMatch[4] <= rightMatch[4], nil
			},
		},
	}

	// This function finds a strategy which recognizes the versions passed to it, then applies that strategy.
	return func(left, right string) (bool, error) {
		if left == right {
			return true, nil
		}

		for _, strategy := range wellKnownStrategies {
			if leftMatch, rightMatch := strategy.match.FindAllStringSubmatch(left, 1), strategy.match.FindAllStringSubmatch(right, 1); len(leftMatch) > 0 && len(rightMatch) > 0 {
				return strategy.handler(leftMatch[0], rightMatch[0])
			}
		}
		return false, fmt.Errorf("Unable to find versioning strategy that could compare %q and %q", left, right)
	}
}()

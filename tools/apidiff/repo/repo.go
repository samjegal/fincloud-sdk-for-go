package repo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type WorkingTree struct {
	dir string
}

func Get(dir string) (wt WorkingTree, err error) {
	orig := dir
	for found := false; !found; {
		var fi []os.FileInfo
		fi, err = ioutil.ReadDir(dir)
		if err != nil {
			return
		}
		// look for the .git directory
		for _, f := range fi {
			if f.Name() == ".git" {
				found = true
				break
			}
		}
		if !found {
			// not the root of the repo, move the parent directory
			i := strings.LastIndexByte(dir, os.PathSeparator)
			if i < 0 {
				err = fmt.Errorf("failed to find repo root under '%s'", orig)
				return
			}
			dir = dir[:i]
		}
	}
	wt.dir = dir
	return
}

func (wt WorkingTree) Branch() (string, error) {
	cmd := exec.Command("git", "branch")
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(output))
	}
	branches := strings.Split(string(output), "\n")
	for _, branch := range branches {
		if branch[0] == '*' {
			return branch[2:], nil
		}
	}
	return "", fmt.Errorf("failed to determine active branch: %s", strings.Join(branches, ","))
}

func (wt WorkingTree) DeleteBranch(branchName string) error {
	cmd := exec.Command("git", "branch", "-d", branchName)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

func (wt WorkingTree) Clone(dest string) (result WorkingTree, err error) {
	cmd := exec.Command("git", "clone", fmt.Sprintf("file://%s", wt.dir), dest)
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = errors.New(string(output))
		return
	}
	result.dir = dest
	return
}

func (wt WorkingTree) Checkout(tree string) error {
	cmd := exec.Command("git", "checkout", tree)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

func (wt WorkingTree) Root() string {
	return wt.dir
}

type CherryCommit struct {
	// Hash is the SHA1 of the commit.
	Hash string

	// Found indicates if the commit was found in the upstream branch.
	Found bool
}

func (wt WorkingTree) Cherry(upstream string) ([]CherryCommit, error) {
	cmd := exec.Command("git", "cherry", upstream)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(output))
	}

	items := strings.Split(string(output), "\n")
	// skip the last entry as it's just an empty string
	commits := make([]CherryCommit, len(items)-1, len(items)-1)
	for i := 0; i < len(items)-1; i++ {
		commits[i].Found = items[i][0] == '-'
		commits[i].Hash = items[i][2:]
	}
	return commits, nil
}

func (wt WorkingTree) CherryPick(commit string) error {
	cmd := exec.Command("git", "cherry-pick", commit)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

func (wt WorkingTree) CreateTag(name string) error {
	cmd := exec.Command("git", "tag", name)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

func (wt WorkingTree) ListTags(pattern string) ([]string, error) {
	cmd := exec.Command("git", "tag", "-l", pattern)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(output))
	}
	if len(output) == 0 {
		return []string{}, nil
	}
	tags := strings.Split(strings.TrimSpace(string(output)), "\n")
	sort.Strings(tags)
	return tags, nil
}

func (wt WorkingTree) Pull(upstream, branch string) error {
	cmd := exec.Command("git", "pull", upstream, branch)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

func (wt WorkingTree) CreateAndCheckout(branch string) error {
	cmd := exec.Command("git", "checkout", "-b", branch)
	cmd.Dir = wt.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

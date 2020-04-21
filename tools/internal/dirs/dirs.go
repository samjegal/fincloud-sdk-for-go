package dirs

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetSubdirs(current string) ([]string, error) {
	children, err := ioutil.ReadDir(current)
	if err != nil {
		return nil, err
	}
	dirs := []string{}
	for _, info := range children {
		if info.IsDir() {
			dirs = append(dirs, info.Name())
		}
	}
	return dirs, nil
}

func DeleteChildDirs(dir string) error {
	children, err := GetSubdirs(dir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, child := range children {
		childPath := filepath.Join(dir, child)
		err = os.RemoveAll(childPath)
		if err != nil {
			return err
		}
	}
	return nil
}

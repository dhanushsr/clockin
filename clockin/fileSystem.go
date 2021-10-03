package clockin

import (
	"os"
	"path"
	"strings"
)

type Entry struct {
	Timestamp string
	Project   string
	Module    string
	Remarks   string
}

type EntryFile struct {
	Entries []Entry
}

func SanitizePath(Path string) (string, error) {
	isAbsolute := path.IsAbs(Path)
	if isAbsolute {
		return Path, nil
	}
	if Path[0] != 126 {
		cwd, err := os.Getwd()
		if err != nil {
			return "", nil
		}
		return appendPathSeperator(path.Clean(path.Join(cwd, Path))), nil
	}
	userHomeDir, err := os.UserHomeDir()
	Path = strings.Replace(Path, "~/", "", -1)
	if err != nil {
		return "", err
	}
	return appendPathSeperator(path.Clean(path.Join(userHomeDir, Path))), nil
}

func appendPathSeperator(Path string) string {
	_, filePath := path.Split(Path)
	if len(filePath) > 0 {
		return Path
	}
	if Path[len(Path)-1] != os.PathSeparator {
		Path = Path + string(os.PathSeparator)
	}
	return Path
}

func DoesPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateIfNotPresent(path string) error {
	doesPathExist, err := DoesPathExist(path)
	if err != nil {
		return err
	}
	if !doesPathExist {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

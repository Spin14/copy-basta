package parse

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type ignorer struct {
	dirs     []string
	patterns []string
}

func newIgnorer(root string, r io.Reader) (*ignorer, error) {
	i := ignorer{}
	if r == nil {
		return &i, nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasSuffix(line, "/") {
			// completely excluded dir
			dir := filepath.Join(root, line)
			i.dirs = append(i.dirs, dir)
		} else {
			// patterns
			pattern := filepath.Join(root, line)
			if _, err := filepath.Match(pattern, ""); err != nil {
				return nil, err
			}

			i.patterns = append(i.patterns, pattern)
		}
	}
	return &i, nil
}

func (i *ignorer) ignore(s string) bool {
	for _, dir := range i.dirs {
		target := s
		for {
			target = filepath.Dir(target)
			if target == "." {
				break
			}
			if target == dir {
				return true
			}
		}
	}

	for _, pattern := range i.patterns {
		matched, err := filepath.Match(pattern, s)
		if err != nil {
			return false
		}
		if matched {
			return matched
		}
	}

	return false
}

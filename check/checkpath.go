// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PathChecker struct {
	err       error
	gopathbin string
}

func (c *PathChecker) Check() (bool, bool) {
	gopath, err := runCmd("go", "env", "GOPATH")
	if err != nil {
		c.err = err
		return false, false
	}

	c.gopathbin = filepath.Clean(filepath.Join(gopath, "bin"))
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	for _, p := range paths {
		if filepath.Clean(p) == c.gopathbin {
			return true, false
		}
	}
	return false, false
}

func (c *PathChecker) Summary() string {
	return fmt.Sprintf("Checking if $PATH contains %q", c.gopathbin)
}

func (c *PathChecker) Resolution() string {
	// TODO(jbd): Add windows specific instructions.
	return fmt.Sprintf(`Add %q to your $PATH.
On Unix systems:
export PATH=$PATH:%v`, c.gopathbin, c.gopathbin)
}

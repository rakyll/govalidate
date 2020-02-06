// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"fmt"
	"os/exec"
	"strings"
)

var supportedGoVersions = []string{ // TODO(jbd): Get the list from golang.org
	"go1.12.16",
	"go1.13.7",
	"go1.14rc1",
}

type GoChecker struct {
	version string
	err     error
}

func (g *GoChecker) Check() (bool, bool) {
	_, err := exec.LookPath("go")
	if err != nil {
		g.err = err
		return false, false
	}
	versionStr, err := runCmd("go", "version")
	if err != nil {
		g.err = err
		return false, false
	}

	vparts := strings.Split(versionStr, " ")
	g.version = vparts[2]
	for _, v := range supportedGoVersions {
		if g.version == v {
			return true, false
		}
	}
	return false, false
}

func (g *GoChecker) Summary() string {
	if g.version == "" {
		return "Go installation"
	}
	return fmt.Sprintf("Go (%v)", g.version)
}

func (g *GoChecker) Resolution() string {
	if g.err != nil {
		return fmt.Sprintf(`Is Go installed? %v.
Visit https://golang.org/dl/ to download Go.`, g.err)
	}

	return fmt.Sprintf(`Your current Go version (%v) is old.
Current Go versions are %v.
Visit https://golang.org/dl/ for a new version.`,
		g.version, strings.Join(supportedGoVersions, ", "))
}

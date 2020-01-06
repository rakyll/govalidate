// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var supportedGoVersions = []string{ // TODO(jbd): Get the list from golang.org
	"go1.12.14",
	"go1.13.5",
	"go1.14beta1",
}

type goChecker struct {
	version string
	err     error
}

func (g *goChecker) check() (bool, bool) {
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

func (g *goChecker) summary() string {
	if g.version == "" {
		return "Go installation"
	}
	return fmt.Sprintf("Go (%v)", g.version)
}

func (g *goChecker) resolution() string {
	if g.err != nil {
		return fmt.Sprintf(`Is Go installed? %v.
Visit https://golang.org/dl/ to download Go.`, g.err)
	}

	return fmt.Sprintf(`Your current Go version (%v) is not supported.
Visit https://golang.org/dl/ for a supported version.`, g.version)
}

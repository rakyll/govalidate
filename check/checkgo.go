// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

const versionsURL = "https://raw.githubusercontent.com/rakyll/govalidate/master/goversion.txt"

type GoChecker struct {
	version             string
	supportedGoVersions []string
	err                 error
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

	req, err := http.Get(versionsURL)
	if err != nil {
		g.err = err
		return false, false
	}
	defer req.Body.Close()

	versions, err := readGoVersions(req.Body)
	if err != nil {
		g.err = err
		return false, false
	}
	g.supportedGoVersions = versions

	vparts := strings.Split(versionStr, " ")
	g.version = vparts[2]
	for _, v := range versions {
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
		g.version, strings.Join(g.supportedGoVersions, ", "))
}

func readGoVersions(r io.Reader) ([]string, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	versions := strings.Split(bodyStr, "\n")
	return versions, nil
}

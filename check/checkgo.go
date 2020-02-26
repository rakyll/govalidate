// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

const downloadsURL = "https://golang.org/dl/?mode=json"

type GoChecker struct {
	version             string   // available after Check is run.
	supportedGoVersions []string // available after Check is run.
	err                 error    // available after Check is run.
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

	versions, err := readGoVersions()
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

func readGoVersions() ([]string, error) {
	req, err := http.Get(downloadsURL)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var versionsJSON []versionJSON
	if err := json.Unmarshal(body, &versionsJSON); err != nil {
		return nil, err
	}

	var versions = make([]string, len(versionsJSON))
	for i, vJSON := range versionsJSON {
		versions[i] = vJSON.Version
	}
	return versions, nil
}

type versionJSON struct {
	Version string `json:"version,omitempty"`
}

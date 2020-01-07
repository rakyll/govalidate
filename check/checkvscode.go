// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"bufio"
	"os/exec"
)

type VSCodeChecker struct{}

func (c *VSCodeChecker) Check() (bool, bool) {
	_, err := exec.LookPath("code")
	if err != nil {
		return false, true // if no VSCode is installed, skip.
	}

	cmd := exec.Command("code", "--list-extensions")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return false, false
	}
	defer out.Close()
	if err := cmd.Start(); err != nil {
		return false, false
	}

	// Return as soon as the extension name appears in the output.
	// Otherwise, command hangs for a while before it finally exists.
	reader := bufio.NewReader(out)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return false, false
		}
		if line == "ms-vscode.Go\n" {
			return true, false
		}
	}
}

func (c *VSCodeChecker) Summary() string {
	return "VSCode Go extension"
}

func (c *VSCodeChecker) Resolution() string {
	return `VSCode Go extension is not installed.
See https://code.visualstudio.com/docs/languages/go to install.`
}

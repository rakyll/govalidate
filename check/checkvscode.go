// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"bufio"
	"os/exec"
	"strings"
)

const goExtensionName = "ms-vscode.go"

type VSCodeChecker struct{}

func (c *VSCodeChecker) Check() (bool, bool) {
	_, err := exec.LookPath("code")
	if err != nil {
		return false, true // if no VSCode is installed, skip.
	}

	ok := isInstalledIn("code", "--list-extensions")
	if !ok {
		// Validate if it is running on a Windows OS.
		// To validate WSL environment.
		if isWindows() {
			return isInstalledIn("wsl", "ls", "~/.vscode-server/extensions"), false
		}
	}
	return true, false
}

func (c *VSCodeChecker) Summary() string {
	return "VSCode Go extension"
}

func (c *VSCodeChecker) Resolution() string {
	return `VSCode Go extension is not installed.
See https://code.visualstudio.com/docs/languages/go to install.`
}

func isInstalledIn(command string, arg ...string) bool {
	cmd := exec.Command(command, arg...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return false
	}
	defer out.Close()
	if err := cmd.Start(); err != nil {
		return false
	}

	reader := bufio.NewReader(out)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return false
		}
		contains := strings.Contains(strings.ToLower(line), GoExtensionName)
		if contains {
			return true
		}
	}
}

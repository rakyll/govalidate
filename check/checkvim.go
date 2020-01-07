// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

type VimChecker struct{}

func (c *VimChecker) Check() (bool, bool) {
	_, err := exec.LookPath("vim")
	if err != nil {
		return false, true // if no vim is installed, skip.
	}
	var ok bool
	homeDir := guessHomeDir()
	filepath.Walk(filepath.Join(homeDir, ".vim"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "vim-go") {
			ok = true
			return io.EOF // no need to look for more
		}
		return nil
	})
	return ok, false
}

func (c *VimChecker) Summary() string {
	return "Vim Go plugin"
}

func (c *VimChecker) Resolution() string {
	return `Vim is installed but the Go plugin is not available.
See https://github.com/fatih/vim-go to install.`
}

func guessHomeDir() string {
	// Prefer $HOME over user.Current due to glibc bug: golang.org/issue/13470
	if v := os.Getenv("HOME"); v != "" {
		return v
	}
	// Else, fall back to user.Current:
	if u, err := user.Current(); err == nil {
		return u.HomeDir
	}
	return ""
}

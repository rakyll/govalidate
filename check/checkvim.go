// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

type VimChecker struct {
	err error
}

func (c *VimChecker) Check() (bool, bool) {
	_, err := exec.LookPath("vim")
	if err != nil {
		return false, true // if no vim is installed, skip.
	}
	var ok bool
	vimDir := filepath.Join(guessHomeDir(), ".vim")
	fi, err := os.Lstat(vimDir)
	if err != nil {
		c.err = err
		return false, false
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		resolved, err := os.Readlink(vimDir)
		if err != nil {
			c.err = err
			return false, false
		}
		vimDir = resolved
	}

	filepath.Walk(vimDir, func(path string, info os.FileInfo, err error) error {
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
	var msg string
	msg = "Vim is installed but cannot determine the Go plugin status.\n"
	if c.err != nil {
		msg += fmt.Sprintf("Error: %v\n", c.err)
	}
	return msg + "See https://github.com/fatih/vim-go to install."
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

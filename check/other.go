// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"bytes"
	"os/exec"
)

func runCmd(cmd string, arg ...string) (string, error) {
	c := exec.Command(cmd, arg...)
	out, err := c.CombinedOutput()
	return string(bytes.TrimSpace(out)), err
}

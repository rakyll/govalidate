// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import (
	"fmt"
	"strings"
)

type CGOChecker struct{}

func (c *CGOChecker) Check() (bool, bool) {
	// Run gcc and check the output instead of looking
	// in the path. On darwin, if gcc is installed but
	// license is not agreed, it shows up a different
	// message.
	gccStr, _ := runCmd("gcc") // will always return an error
	return strings.Contains(gccStr, "no input files"), false
}

func (c *CGOChecker) Summary() string {
	return "Checking gcc for CGO support"
}

func (c *CGOChecker) Resolution() string {
	return fmt.Sprintf(`If you are going to use CGO, install a C compiler.
- On macOS, install XCode and run "xcode-select --install" to install command line tools.
  Then, you may need to accept the license by running "xcodebuild -license".
- On Windows and Linux, install gcc. See https://gcc.gnu.org/install/binaries.html.
If you are not using CGO or using a different C compiler, ignore this message.`)
}

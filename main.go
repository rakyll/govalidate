// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command godoctor checks whether the current system
// is properly configured for Go development.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rakyll/govalidate/check"
)

type checker interface {
	// Check validates a condition and returns ok=true
	// if condition is satisfied. Return skip=true
	// if you don't want the results to be printed.
	Check() (ok bool, skip bool)
	Summary() string
	Resolution() string
}

var (
	ignoreCGO     bool
	ignoreEditors bool
)

func main() {
	flag.BoolVar(&ignoreCGO, "ignore-cgo", false, "")
	flag.BoolVar(&ignoreEditors, "ignore-editors", false, "")
	flag.Parse()

	var exit int
	// TODO(jbd): Check operating system requirements.
	// See https://github.com/golang/go/wiki/MinimumRequirements for
	// a more comprehensive list.
	checks := []checker{
		&check.GoChecker{},   // checks go and go version
		&check.PathChecker{}, // checks $GOPATH/bin is in $PATH
	}
	for _, c := range checks {
		exit += runCheck(false, c)
	}
	// Optional checks.
	var optionals []checker
	if !ignoreCGO {
		optionals = append(optionals, &check.CGOChecker{})
	}
	if !ignoreEditors {
		// TODO(jbd): Add Goland.
		optionals = append(optionals,
			&check.VimChecker{},
			&check.VSCodeChecker{},
		)
	}
	for _, c := range optionals {
		exit += runCheck(true, c)
	}

	if exit > 0 {
		os.Exit(1)
	}
}

func runCheck(optional bool, c checker) int {
	var exit int

	ok, skip := c.Check()
	if skip {
		return exit
	}
	if ok {
		color.New(color.FgHiGreen).Print("[✔]")
	} else {
		if !optional {
			exit = 1
			color.New(color.FgRed).Print("[✗]")
		} else {
			color.New(color.FgYellow).Print("[!]")
		}
	}
	fmt.Print(" ")
	fmt.Println(c.Summary())
	if !ok {
		printWithTabs(c.Resolution())
	}
	return exit
}

func printWithTabs(msg string) {
	lines := strings.Split(msg, "\n")
	for _, l := range lines {
		fmt.Printf("    %v\n", l)
	}
}

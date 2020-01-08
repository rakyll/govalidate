# govalidate

[![CircleCI](https://circleci.com/gh/rakyll/govalidate.svg?style=svg&circle-token=8ea1ac2ae17cbac9a5505d875261eb74061f8404)](https://circleci.com/gh/rakyll/govalidate)

Validates your Go installation and dependencies.

* Checks the Go installation and version.
* Checks if the PATH is correctly configured.
* Checks CGO dependencies are installed.
* Checks the plugin support for available editors.

```
$ govalidate
[✔] Go (go1.13.5)
[✗] Checking if $PATH contains "/Users/jbd/go/bin"
    Add "/Users/jbd/go/bin" to your $PATH.
    On Unix systems:
    export PATH=$PATH:/Users/jbd/go/bin
[✔] Checking gcc for CGO support
[✔] Vim Go plugin
[!] VSCode Go extension
    VSCode Go extension is not installed.
    See https://code.visualstudio.com/docs/languages/go to install.
```

## Installation

```
$ go get -u github.com/rakyll/govalidate
```

Or download one of the binaries and run:

* Linux 64-bit: https://storage.googleapis.com/jbd-releases/govalidate_linux_amd64
* macOS 64-bit: https://storage.googleapis.com/jbd-releases/govalidate_darwin_amd64
* Windows 64-bit: https://storage.googleapis.com/jbd-releases/govalidate_windows_amd64

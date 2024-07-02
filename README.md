# Gordinal

The module allows you to execute various commands by pressing hotkeys.

## Requirements

#### For All:

* [Golang](https://go.dev/)
* [GCC](https://gcc.gnu.org/)

#### For MacOS:

* Xcode Command Line Tools

```shell
xcode-select --install
```

#### For Windows:

* [MinGW](https://sourceforge.net/projects/mingw-w64/files)

> Download the Mingw, then set system environment variables `C:\mingw64\bin` to the Path.  
> [Set environment variables to run GCC from command line.](https://www.youtube.com/results?search_query=Set+environment+variables+to+run+GCC+from+command+line)

#### For Linux:

Ubuntu:

```shell
sudo apt install \
  gcc \
  libc6-dev \
  xcb \
  libxcb-xkb-dev \
  x11-xkb-utils \
  libx11-xcb-dev \
  libxkbcommon-x11-dev \
  libxkbcommon-dev
```

Fedora:

```shell
sudo dnf install \
  libXtst-devel \
  libxkbcommon-devel \
  libxkbcommon-x11-devel \
  xorg-x11-xkb-utils-devel
```

## Install

```shell
go get github.com/Gwinkamp/gordinal
```

## Usage

#### From golang code:

```go
package main

import (
  "fmt"
  "github.com/Gwinkamp/gordinal"
)


func main() {
	g := gordinal.New()

	g.Register(
		"simple",
		[]string{"ctrl", "shift", "1"},
		func() error {
			fmt.Println("the command is executed here")
			return nil
		},
	)

	g.Run()
}
```

#### From config file

You can describe the keyboard shortcuts and commands in the configuration file.  
An example of such a file:  

```yaml
hooks:
  - name: simple
    keys:
      - ctrl
      - shift
      - 1
    command: echo
    args:
      - Hello world
logging:
  level: debug
  output: stdout
```

* `hooks` contains a list of instructions for running a command by hotkey [required]
  * `name` - unique name required to identify the command [required]
  * `keys` - list of hotkeys [required]
  * `command` - shell command or path to exec file [required]
  * `args` - list of arguments for command [optional]
* `logging` contains setting for logging [optional]
  * `level` - string with the logging level [optional, default:info]
  * `output` - string indicating where to output logs (filepath or stdout) [optional, default:stdout]

To launch:

```shell
go run cmd/main/main.go --path {path-to-config}
```

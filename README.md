# Gordinal

The module allows you to execute various commands by pressing hotkeys.

# Install

```shell
go get github.com/Gwinkamp/gordinal
```

# Usage

### From golang code:

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

### From config file

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

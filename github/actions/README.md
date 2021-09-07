# GitHub / Actions
Installation: `go get github.com/whiteducksoftware/golang-utilities/github/actions`

## Types & Functions
```go
package github

// Types
type GitHub struct {}   // represents the inputs which github provides us on default
type Message struct {}  // is a struct used for printing warning/error messages

// Functions
func SetOutput(name string, value string) {} // can be used to set the outputs of your action
func WriteDebug(message string) {}      // is used to write debug message to the github actions log
func WriteWarning(message Message) {}   // is used to write warning message to the github actions log
func WriteError(message Message) {}     // is used to write error message to the github actions log

func (g *GitHub) Load() error {}        // parses the environment vars and reads github options

func LoadOptions() (GitHub, error) {}   // parses the environment vars and reads github options 
                                        // Deprecated: Use GitHub.Load instead.
```
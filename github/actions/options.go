// Package actions contains utilities for working with github actions
package actions

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// GitHub represents the inputs which github provides us on default
type GitHub struct {
	Workflow           string `env:"GITHUB_WORKFLOW"`
	RunID              uint16 `env:"GITHUB_RUN_ID"`
	JobID              string `env:"GITHUB_JOB"`
	Action             string `env:"GITHUB_ACTION"`
	Actor              string `env:"GITHUB_ACTOR"`
	Repository         string `env:"GITHUB_REPOSITORY"`
	Commit             string `env:"GITHUB_SHA"`
	EventName          string `env:"GITHUB_EVENT_NAME"`
	EventPath          string `env:"GITHUB_EVENT_PATH"`
	Ref                string `env:"GITHUB_REF"`
	RunningAsAction    bool   `env:"GITHUB_ACTIONS" envDefault:"false"`
	OperatingSystem    string `env:"RUNNER_OS"`
	TemporaryDirectory string `env:"RUNNER_TEMP"`
}

// Load parses the environment vars and reads github options
func (g *GitHub) Load() error {
	if err := env.Parse(g); err != nil {
		return fmt.Errorf("failed to parse github environments: %s", err)
	}

	return nil
}

// LoadOptions parses the environment vars and reads github options
// Deprecated: Use GitHub.Load instead.
func LoadOptions() (GitHub, error) {
	github := GitHub{}
	if err := github.Load(); err != nil {
		return GitHub{}, err
	}

	return github, nil
}

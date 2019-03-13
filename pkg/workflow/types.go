package workflow

import (
	"github.com/puppetlabs/nebula/pkg/errors"
	"github.com/puppetlabs/nebula/pkg/workflow/runner"
	"gopkg.in/yaml.v2"
)

type Workflow struct {
	Version   string     `yaml:"version"`
	Name      string     `yaml:"name"`
	Variables []Variable `yaml:"variables"`
	Actions   []Action   `yaml:"actions"`
	Stages    []Stage    `yaml:"stages"`
}

func (w Workflow) Stage(name string) (*Stage, errors.Error) {
	for _, stage := range w.Stages {
		if stage.Name == name {
			return &stage, nil
		}
	}

	return nil, errors.NewWorkflowStageDoesNotExist(name)
}

type Trigger struct {
	Action string `yaml:"action"`
	Branch string `yaml:"branch"`
}

type Stage struct {
	Name        string    `yaml:"name"`
	ActionNames []string  `yaml:"actions"`
	Trigger     []Trigger `yaml:"trigger"`

	actions []Action
}

func (s *Stage) Actions() []Action {
	return s.actions
}

func (s *Stage) AddAction(a Action) {
	s.actions = append(s.actions, a)
}

type Action struct {
	Name       string                      `yaml:"name"`
	Kind       string                      `yaml:"kind"`
	ResourceID string                      `yaml:"resourceID"`
	Spec       map[interface{}]interface{} `yaml:"spec"`

	loadedRunner runner.ActionRunner
}

func (a Action) Runner() runner.ActionRunner {
	return a.loadedRunner
}

func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}

	if err := unmarshal(&raw); err != nil {
		return err
	}

	name, ok := raw["name"].(string)
	if !ok {
		return errors.NewWorkflowActionDecodeError("`name` was not a string")
	}
	a.Name = name

	kind, ok := raw["kind"].(string)
	if !ok {
		return errors.NewWorkflowActionDecodeError("`kind` was not a string")
	}
	a.Kind = kind

	if rawResourceID, ok := raw["resourceID"]; ok {
		resourceID, ok := rawResourceID.(string)
		if !ok {
			return errors.NewWorkflowActionDecodeError("`resourceID` was not a string")
		}
		a.ResourceID = resourceID
	}

	if _, ok := raw["spec"]; ok {
		a.Spec = raw["spec"].(map[interface{}]interface{})
	}

	r, err := runner.NewRunner(runner.RunnerKind(a.Kind))
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(raw)
	if err != nil {
		return err
	}

	if err := r.Decoder().Decode(b); err != nil {
		return err
	}

	a.loadedRunner = r

	return nil
}

type Variable struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
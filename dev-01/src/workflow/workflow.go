package workflow

import (
	"fmt"

	"../entity/agent"
)

type Input struct{}

type Output struct{}

type TaskType string

const (
	Implementation = TaskType("implementation")
)

type ClassificatorInterface interface {
	Evaluate(input Input) (TaskType, error)
}

type Classificator struct{}

func (c *Classificator) Evaluate(input Input) (TaskType, error) {
	return "", nil
}

type PromptInterface interface {
	Run(input Input) (Output, error)
}

//fazer um map de string com agents e ir usando as chaves das strings, que seriam os kinds, para encontrar o agente certo?

type Workflow struct {
	Classificator Classificator
	AgentsMap     map[TaskType]agent.AgentInterface
	//logs
}

func (w *Workflow) Run(input Input) (Output, error) {
	taskKind, err := w.Classificator.Evaluate(input)
	if err != nil {
		return Output{}, err
	}
	ag, ok := w.AgentsMap[taskKind]
	if !ok {
		return Output{}, nil
	}
	res, err := ag.Run(agent.AgentInput{})
	if err != nil {
		return Output{}, nil
	}
	fmt.Println(res)
	return Output{}, nil
}

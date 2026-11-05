package agent

type AgentInput struct {
	ex string
}

type AgentOutput struct{}

type AgentInterface interface {
	Run(input AgentInput) (AgentOutput, error)
}

type Agent01 struct{}

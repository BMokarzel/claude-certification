package tool

type ToolInputInterface interface{}

type ToolOutputInterface interface{}

type ToolInput struct{}

type ToolOutput struct{}

func ToolExec(input ToolInputInterface) (ToolOutputInterface, error) {
	return ToolOutput{}, nil
}

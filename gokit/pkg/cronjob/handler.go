package cronjob

import "context"

type InputParams any
type Result any

type ParameterHandler func(context.Context) (InputParams, error)
type ExecutionHandler func(context.Context, InputParams) (Result, error)
type ResultHandler func(context.Context, Result, InputParams) error
type RetireHandler func(context.Context, *Cronjob) error

type Handler struct {
	parameter ParameterHandler
	execute   ExecutionHandler
	result    ResultHandler
	retire    RetireHandler
}

func NewHandler(
	parameter ParameterHandler, execute ExecutionHandler, result ResultHandler, retire RetireHandler,
) *Handler {
	return &Handler{
		parameter: parameter,
		execute:   execute,
		result:    result,
		retire:    retire,
	}
}

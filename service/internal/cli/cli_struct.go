package cli

import "translate-remember-repeat/internal/domain"

type cliStruct struct {
	enteringMessage string
	casualMessage   string
}

func New() *cliStruct {
	enteringMessage := domain.EnteringMessage
	casualMessage := domain.CasualMessage
	return &cliStruct{
		enteringMessage: enteringMessage,
		casualMessage:   casualMessage,
	}
}

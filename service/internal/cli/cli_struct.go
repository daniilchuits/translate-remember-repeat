package cli

import "translate-remember-repeat/internal/domain"

type cliStruct struct {
	enteringMessage  string
	casualMessage    string
	testAnnouncement string
}

func New() *cliStruct {
	enteringMessage := domain.EnteringMessage
	casualMessage := domain.CasualMessage
	testAnnouncement := domain.TestAnnouncement
	return &cliStruct{
		enteringMessage:  enteringMessage,
		casualMessage:    casualMessage,
		testAnnouncement: testAnnouncement,
	}
}

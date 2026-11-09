package cli

import "translate-remember-repeat/internal/domain"

type cliStruct struct {
	enteringMessage  string
	casualMessage    string
	testAnnouncement string
	zeroSentences    string
	lenSentences     string
	correct          string
	writeTranslation string
	wrongAnswear     string
	separators       string
	separatorsTest   string
	exit             string
}

func New() *cliStruct {
	enteringMessage := domain.EnteringMessage
	casualMessage := domain.CasualMessage
	testAnnouncement := domain.TestAnnouncement
	zero := domain.ZeroSentences
	lSent := domain.LenSentences
	correct := domain.Correct
	writeTranslation := domain.Translate
	wrongAnswear := domain.WrongAnswear
	separators := domain.Separators
	separatorsTest := domain.SeparatorsTest
	exit := domain.Exit
	return &cliStruct{
		enteringMessage:  enteringMessage,
		casualMessage:    casualMessage,
		testAnnouncement: testAnnouncement,
		zeroSentences:    zero,
		lenSentences:     lSent,
		correct:          correct,
		writeTranslation: writeTranslation,
		wrongAnswear:     wrongAnswear,
		separators:       separators,
		separatorsTest:   separatorsTest,
		exit:             exit,
	}
}

package cli

import (
	"fmt"
	"translate-remember-repeat/internal/domain"
)

func (cli *cliStruct) LogResponse(sentence string, resp *domain.ResponseOfDeepL) {

	fmt.Printf(
		"Translation for %s is %s, detached language is %s\n",
		sentence, resp.Text(), resp.DetechedSourceLanguage(),
	)
}

package cli

import "log"

func (cli *cliStruct) Panic(msg string) {
	log.Panic(msg)
}

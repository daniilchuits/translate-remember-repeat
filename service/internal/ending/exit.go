package ending

import (
	"context"
	"log"
)

type cliWriter interface {
	Exit(ctx context.Context) error
	Separate(ctx context.Context) error
}

type endingStruct struct {
	cliWriter
}

func New(cliWriter cliWriter) *endingStruct {
	return &endingStruct{cliWriter: cliWriter}
}

func (end *endingStruct) Exec(ctx context.Context) {
	if err := end.Separate(ctx); err != nil {
		log.Panicf("Exit not finished: %s\n", err.Error())
	}
	if err := end.Exit(ctx); err != nil {
		log.Panicf("Exit not finished: %s\n", err.Error())
	}
	if err := end.Separate(ctx); err != nil {
		log.Panicf("Exit not finished: %s\n", err.Error())
	}
}

package cli

import (
	"fmt"
)

func (c *cliStruct) WriteEnteringMessage() {
	fmt.Print(c.enteringMessage)
}

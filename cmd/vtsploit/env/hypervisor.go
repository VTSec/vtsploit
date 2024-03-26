package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
	"vtsploit/env"
)

var Hypervisor = &cli.Command{
	Name:    "hypervisor",
	Aliases: []string{"r"},
	Usage:   "show hypervisor type",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		result := env.Hypervisor()
		fmt.Println(printer.Printer.Print(result))
		return
	},
}

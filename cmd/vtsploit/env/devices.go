package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
	"vtsploit/env"
)

var Devices = &cli.Command{
	Name:    "devices",
	Aliases: []string{"d"},
	Usage:   "show device list",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("")
		result := env.Devices()
		fmt.Println(printer.Printer.Print(result))
		return
	},
}

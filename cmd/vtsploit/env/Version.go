package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/urfave/cli/v2"
	"vtsploit/env"
)

var Version = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "show qemu version info",
	Action: func(context *cli.Context) (err error) {
		log.Logger.Debug("xxx")
		result := env.Version()
		fmt.Println(printer.Printer.Print(result))
		return
	},
}

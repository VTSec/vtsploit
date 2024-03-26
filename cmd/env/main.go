
package main

import (
	"vtsploit/cmd/vtsploit/env"
	"vtsploit/pkg/app"
	"os"
)

const (
	name = `vtsploit/env`
)

func main() {
	sploit := app.Command2App(env.Command)
	sploit.Name = name
	app.InstallGlobalFlags(sploit)
	err := sploit.Run(os.Args)
	if err != nil {
		awesome_error.CheckFatal(err)
	}
}
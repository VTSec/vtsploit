package main

import (
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/ctrsploit/sploit-spec/pkg/version"
	"github.com/urfave/cli/v2"
	"os"
	"vtsploit/cmd/vtsploit/auto"
	"vtsploit/cmd/vtsploit/checksec"
	"vtsploit/cmd/vtsploit/env"
)

const usage = `An example sploit tool follows sploit-spec`

func init() {
	version.ProductName = "vtsploit"
}

func main() {
	sploit := &cli.App{
		Name:  "vtsploit",
		Usage: usage,
		Commands: []*cli.Command{
			auto.Command,
			env.Command,
			checksec.Command,
			version.Command,
		},
	}
	app.InstallGlobalFlags(sploit)
	_ = sploit.Run(os.Args)
}

package checksec

import (
	"github.com/ctrsploit/sploit-spec/pkg/vul"
	vul2 "vtsploit/vul"
)

const (
	CommandNameAuto = "auto"
)

var (
	Auto = &cli.Command{
		Name:    CommandNameAuto,
		Usage:   "auto",
		Aliases: []string{"a"},
		Action: func(context *cli.Context) (err error) {
			vulnerabilities := vul.Vulnerabilities{
				vul2.CVE_2099_9999_v1,
			}
			err = vulnerabilities.Check()
			if err != nil {
				return
			}
			vulnerabilities.Output()
			return
		},
	}
)

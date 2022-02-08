package cmd

import (
	_ "github.com/jakubenglicky/enu/cmd/crontab"
	"github.com/jakubenglicky/enu/cmd/root"
	_ "github.com/jakubenglicky/enu/cmd/version"
)

func Excecute() {
	root.RootCmd.Execute()
}

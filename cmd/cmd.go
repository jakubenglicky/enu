package cmd

import (
	_ "github.com/jakubenglicky/enu/cmd/crontab"
	"github.com/jakubenglicky/enu/cmd/root"
)

func Excecute() {
	root.RootCmd.Execute()
}

package crontab

import (
	"fmt"

	"github.com/jakubenglicky/enu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "crontab",
	Short: "Generate CronJobs YAML from Unix crontab",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crontab Generator")
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}

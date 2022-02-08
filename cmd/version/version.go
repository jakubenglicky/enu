package version

import (
	"fmt"

	"github.com/jakubenglicky/enu/cmd/root"
	"github.com/jakubenglicky/enu/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "EnUtils installed version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: " + version.Version)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}

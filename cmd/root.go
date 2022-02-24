package cmd

import (
	_ "embed"
	"fmt"
	"os"

	whris "github.com/harakeishi/whris/whris"
	"github.com/spf13/cobra"
)

//go:embed version.txt
var Version string

var rootCmd = &cobra.Command{
	Use:   "whris",
	Short: "`whris` is Displays management information for IPs associated with the domain.",
	Long:  `"whris" outputs the target domain and IP from the input domain, as well as the administrator information for that IP (administrator name, network name, range of IPs to be managed, and country name).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if version {
			fmt.Println(Version)
			return nil
		}
		domain := args[0]
		v, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := whris.Resolve(domain, v); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "echo version")
}

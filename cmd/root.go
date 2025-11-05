package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	project string
)

var rootCmd = &cobra.Command{
	Use:   "iclone",
	Short: "clone and install deps of project",
	Long: `
		It clones provided repo using git and install dependencies according to project type. eg. npm, pnpm, go, rust etc..
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(project)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&project, "project", "p", "AUTO", "Help message for toggle")
	rootCmd.CompletionOptions.DisableDefaultCmd = false
}

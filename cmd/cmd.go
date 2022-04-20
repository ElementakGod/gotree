package cmd

import (
	"github.com/ElementakGod/gotree/pkg"
	"github.com/spf13/cobra"
)

var name, path string

func init() {
	RootCmd.Flags().StringVarP(&name, "name", "n", "", "project name")
	RootCmd.Flags().StringVarP(&path, "path", "p", "", "project path(can be blank)")
}

var RootCmd = &cobra.Command{
	Use:   "gotree",
	Short: "A tool for golang project tree",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(name) == 0 {
			cmd.Help()
			return nil
		}
		return pkg.NewProjectTree(name, path).Setup()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}

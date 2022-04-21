package cmd

import (
	"github.com/ElementakGod/gotree/pkg"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var name, path string
var subDir bool

func init() {
	RootCmd.Flags().StringVarP(&name, "name", "n", "", "project name")
	RootCmd.Flags().StringVarP(&path, "path", "p", "", "project path(can be blank)")
	RootCmd.Flags().BoolVarP(&subDir, "subDir", "s", false, "only create subdirectory")
}

var RootCmd = &cobra.Command{
	Use:   "gotree",
	Short: "A tool for golang project tree",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(name) == 0 {
			cmd.Help()
			return nil
		}
		return pkg.NewProjectTree(name, path, subDir).Setup()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

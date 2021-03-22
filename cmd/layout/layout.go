package main

import (
	"github.com/go-creed/project-layout/pkg"
	"github.com/spf13/cobra"
)

var (
	cmdPackageName string
	cmdPackagePath string
	cmdWithGit     bool
	cmdWithGoMod   bool
)

func main() {
	cmdLayOut := &cobra.Command{
		Use: "init",
		Run: func(cmd *cobra.Command, args []string) {
			g := pkg.NewGenerate(cmdPackageName, cmdPackagePath, cmdWithGit, cmdWithGoMod)
			g.Output()
		},
	}
	cmdLayOut.Flags().StringVarP(&cmdPackageName, "packageName", "n", "", "create package name")
	cmdLayOut.Flags().StringVarP(&cmdPackagePath, "packagePath", "p", "", "package path,can use abs or rel path")
	cmdLayOut.Flags().BoolVarP(&cmdWithGit, "withGit", "g", true, "git init")
	cmdLayOut.Flags().BoolVarP(&cmdWithGoMod, "withGoMod", "m", true, "go mod init")

	err := cmdLayOut.Execute()
	if err != nil {
		panic(err)
	}
}

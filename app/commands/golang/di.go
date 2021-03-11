package golang

import "github.com/spf13/cobra"

//NewRoot root command for golang
func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "go",
		Short: "golang",
	}

	cmd := New()
	root.AddCommand(cmd.GoInstall())
	root.AddCommand(cmd.Envs())

	return root
}

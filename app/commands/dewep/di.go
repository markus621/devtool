package dewep

import "github.com/spf13/cobra"

var dewepNewProject = ""

//NewRoot root command for golang
func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "dewep",
		Short: "helper for github.com/deweppro",
	}

	root.AddCommand(NewService())
	root.AddCommand(NewProject())

	return root
}

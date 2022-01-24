package commands

import "github.com/spf13/cobra"

func createProjectCommand() *cobra.Command {
	create := &cobra.Command{
		Use:   "create",
		Short: "Create new project",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	return create
}

func projectCommand() *cobra.Command {
	project := &cobra.Command{
		Use:   "project",
		Short: "Project commands",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	project.AddCommand(createProjectCommand())

	return project
}

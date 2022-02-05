package commands

import (
	"fmt"
	"log"

	"github.com/dan-lovelace/wink/common"
	winkDB "github.com/dan-lovelace/wink/db"
	"github.com/spf13/cobra"
)

func createProjectCommand(w *common.Wink) *cobra.Command {
	create := &cobra.Command{
		Use:   "create <NAME>",
		Short: "Create new project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Println("Creating project", name)

			db := winkDB.GetDB(w.Context)
			defer db.Close()

			stmt, err := db.Prepare("INSERT INTO project(user_id, name) VALUES((SELECT id FROM user WHERE username = 'default'), ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			res, err := stmt.Exec(name)
			if err != nil {
				log.Fatal(err)
			}

			if _, err := res.LastInsertId(); err != nil {
				log.Fatal(err)
			}
		},
	}

	return create
}

func getProjectsCommand(w *common.Wink) *cobra.Command {
	list := &cobra.Command{
		Use:     "list",
		Short:   "List all projects",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			db := winkDB.GetDB(w.Context)
			defer db.Close()

			stmt, err := db.Prepare("SELECT name FROM project;")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			rows, err := stmt.Query()
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			var names []string
			for rows.Next() {
				var name string
				err := rows.Scan(&name)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(name)
				names = append(names, name)
			}

			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return list
}

func projectCommand(w *common.Wink) *cobra.Command {
	project := &cobra.Command{
		Use:   "project",
		Short: "Project commands",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	project.AddCommand(createProjectCommand(w))
	project.AddCommand(getProjectsCommand(w))

	return project
}

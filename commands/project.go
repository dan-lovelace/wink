package commands

import (
	"fmt"
	"log"

	"github.com/dan-lovelace/wink/common"
	"github.com/dan-lovelace/wink/configs"
	winkDB "github.com/dan-lovelace/wink/db"
	"github.com/spf13/cobra"
)

const currentProjectKey string = "CURRENT_PROJECT"

func createProjectCommand(w *common.Wink) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <NAME>",
		Short: "Create new project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Println("Creating project", name)

			db := winkDB.GetDB(w)
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

	return cmd
}

func getCurrentProjectCommand(w *common.Wink) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display the current project",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			val, err := configs.GetEnv(currentProjectKey)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(val)
		},
	}

	return cmd
}

func getProjectsCommand(w *common.Wink) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all projects",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			db := winkDB.GetDB(w)
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

	return cmd
}

func setProjectCommand(w *common.Wink) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <NAME>",
		Short: "Set the current project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			db := winkDB.GetDB(w)
			defer db.Close()

			stmt, err := db.Prepare("SELECT name FROM project WHERE name = ?;")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			rows, err := stmt.Query(args[0])
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			var ret []string
			for rows.Next() {
				var name string
				err := rows.Scan(&name)
				if err != nil {
					log.Fatal(err)
				}

				ret = append(ret, name)
			}

			if len(ret) < 1 {
				log.Fatal("Project does not exist")
			}

			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}

			configs.SetEnv(currentProjectKey, ret[0])
		},
	}

	return cmd
}

func projectCommand(w *common.Wink) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Project commands",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(createProjectCommand(w))
	cmd.AddCommand(getCurrentProjectCommand(w))
	cmd.AddCommand(getProjectsCommand(w))
	cmd.AddCommand(setProjectCommand(w))

	return cmd
}

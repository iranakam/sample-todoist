/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
}

var projectsListCmd = &cobra.Command{
	Use:           "list",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.Req("", nil)
		if err != nil {
			return err
		}
		cmd.Println("ProjectID\tName\tParentID")
		for _, project := range res.Projects {
			cmd.Println(
				strconv.Itoa(project.ID) +
					"\t" + project.Name +
					"\t" +
					strconv.Itoa(project.ParentID),
			)
		}
		return nil
	},
}

var projectsAddCmd = &cobra.Command{
	Use:           "add",
	SilenceErrors: true,
	Short:         "A brief description of your command",
	Long:          `A longer description that spans multiple lines and likely contains examples`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("project_add", mapArgs)
		return err
	},
}

var projectsUpdateCmd = &cobra.Command{
	Use:           "update",
	SilenceErrors: true,
	Short:         "A brief description of your command",
	Long:          `A longer description that spans multiple lines and likely contains examples`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("project_update", mapArgs)
		return err
	},
}

var projectsMoveCmd = &cobra.Command{
	Use:           "move",
	SilenceErrors: true,
	Short:         "A brief description of your command",
	Long:          `A longer description that spans multiple lines and likely contains examples`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("project_move", mapArgs)
		return err
	},
}

var projectsDeleteCmd = &cobra.Command{
	Use:           "delete",
	SilenceErrors: true,
	Short:         "A brief description of your command",
	Long:          `A longer description that spans multiple lines and likely contains examples`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("project_delete", mapArgs)
		return err
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsAddCmd,
		projectsUpdateCmd,
		projectsMoveCmd,
		projectsDeleteCmd,
	)
	projectsAddCmdInit()
	projectsUpdateCmdInit()
	projectsMoveCmdInit()
	projectsDeleteCmdInit()
}

func projectsAddCmdInit() {
	projectsAddCmd.PersistentFlags().StringVar(&structArgs.Name, "name", "", Messages["project"]["name"])
	projectsAddCmd.PersistentFlags().IntVar(&structArgs.Color, "color", 0, Messages["project"]["color"])
	projectsAddCmd.PersistentFlags().IntVar(&structArgs.ParentID, "parent-id", 0, Messages["project"]["parent-id"])
	projectsAddCmd.PersistentFlags().IntVar(&structArgs.IsFavorite, "is-favorite", 0, Messages["project"]["is-favorite"])
	projectsAddCmd.PersistentFlags().IntVar(&structArgs.ChildOrder, "child-order", 0, Messages["project"]["child-order"])
	projectsAddCmd.MarkPersistentFlagRequired("name")
}

func projectsUpdateCmdInit() {
	projectsUpdateCmd.PersistentFlags().IntVar(&structArgs.ID, "id", 0, Messages["project"]["id"])
	projectsUpdateCmd.PersistentFlags().StringVar(&structArgs.Name, "name", "", Messages["project"]["name"])
	projectsUpdateCmd.PersistentFlags().IntVar(&structArgs.Color, "color", 0, Messages["project"]["color"])
	projectsUpdateCmd.PersistentFlags().IntVar(&structArgs.IsFavorite, "is-favorite", 0, Messages["project"]["is-favorite"])
	projectsUpdateCmd.PersistentFlags().IntVar(&structArgs.Collapsed, "collapsed", 0, Messages["project"]["collapsed"])
	projectsUpdateCmd.MarkPersistentFlagRequired("id")
}

func projectsMoveCmdInit() {
	projectsMoveCmd.PersistentFlags().IntVar(&structArgs.ID, "id", 0, Messages["project"]["id"])
	projectsMoveCmd.PersistentFlags().IntVar(&structArgs.ParentID, "parent-id", 0, Messages["project"]["parent-id"])
	projectsMoveCmd.MarkPersistentFlagRequired("id")
	projectsMoveCmd.MarkPersistentFlagRequired("parent-id")
}

func projectsDeleteCmdInit() {
	projectsDeleteCmd.PersistentFlags().IntVar(&structArgs.ID, "id", 0, Messages["project"]["id"])
	projectsDeleteCmd.MarkPersistentFlagRequired("id")
}

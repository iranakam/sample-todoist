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

// itemsCmd represents the items command
var itemsCmd = &cobra.Command{
	Use:           "items",
	SilenceErrors: false,
	Short:         "Get, Add, Update, Delete, Complete items of todoist.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var itemsListCmd = &cobra.Command{
	Use:           "list",
	SilenceErrors: false,
	Short:         "Get items of your todoist.",
	Long:          `A longer description that spans multiple lines and likely contains examples`,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.Req("", nil)
		if err != nil {
			return err
		}
		cmd.Println("ProjectID\tItemID\tContent\tDueDate")
		for _, item := range res.Items {
			cmd.Println(
				strconv.Itoa(item.ProjectID) +
					"\t" +
					strconv.Itoa(item.ID) +
					"\t" +
					item.Content +
					"\t" +
					item.Due.Date,
			)
		}
		//output.Flush()
		return err
	},
}

var itemsAddCmd = &cobra.Command{
	Use:           "add",
	SilenceErrors: false,
	Short:         "Add item on your todoist.",
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("item_add", mapArgs)
		return err
	},
}

var itemsUpdateCmd = &cobra.Command{
	Use:           "update",
	SilenceErrors: false,
	Short:         "Update specified item.",
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("item_update", mapArgs)
		return err
	},
}

var itemsMoveCmd = &cobra.Command{
	Use:           "move",
	SilenceErrors: false,
	Short:         "Move item to you specified project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("item_move", mapArgs)
		return err
	},
}

var itemsDeleteCmd = &cobra.Command{
	Use:           "delete",
	SilenceErrors: false,
	Short:         "Delete specified item.",
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("item_delete", mapArgs)
		return err
	},
}

var itemsCompleteCmd = &cobra.Command{
	Use:           "complete",
	SilenceErrors: false,
	Short:         "Complete specified item.",
	RunE: func(cmd *cobra.Command, args []string) error {
		mapArgs := structToMap(structArgs)
		_, err := client.Req("item_complete", mapArgs)
		return err
	},
}

func init() {
	rootCmd.AddCommand(itemsCmd)
	itemsCmd.AddCommand(
		itemsListCmd,
		itemsUpdateCmd,
		itemsAddCmd,
		itemsMoveCmd,
		itemsDeleteCmd,
		itemsCompleteCmd,
	)
	itemsAddCmdSetting()
	itemsUpdateCmdSetting()
	itemsMoveCmdSetting()
	itemsDeleteCmdSetting()
	itemsCompleteCmdSetting()
}

func itemsAddCmdSetting() {
	itemsAddCmd.PersistentFlags().StringVarP(&structArgs.Content, "content", "c", "", Messages["item"]["content"])
	itemsAddCmd.PersistentFlags().IntVarP(&structArgs.ProjectID, "project-id", "p", 0, Messages["item"]["project-id"])
	itemsAddCmd.PersistentFlags().IntVarP(&structArgs.Priority, "priority", "r", 0, Messages["item"]["priority"])
	itemsAddCmd.PersistentFlags().IntVarP(&structArgs.ParentID, "parent-id", "P", 0, Messages["item"]["parent-id"])
	itemsAddCmd.PersistentFlags().IntVar(&structArgs.ChildOrder, "child-order", 0, Messages["item"]["child-order"])
	itemsAddCmd.PersistentFlags().IntVar(&structArgs.DayOrder, "day-order", 0, Messages["item"]["day-order"])
	itemsAddCmd.PersistentFlags().IntVar(&structArgs.Collapsed, "collapsed", 0, Messages["item"]["collapsed"])
	itemsAddCmd.PersistentFlags().IntVar(&structArgs.AssignedByUID, "assigned-by-uid", 0, Messages["item"]["assigned-by-uid"])
	itemsAddCmd.PersistentFlags().BoolVar(&structArgs.AutoReminder, "auto-reminder", false, Messages["item"]["auto-reminder"])
	itemsAddCmd.PersistentFlags().BoolVar(&structArgs.AutoParseLabels, "auto-parse-labels", false, Messages["item"]["auto-parse-labels"])
	itemsAddCmd.MarkPersistentFlagRequired("content")
}

func itemsUpdateCmdSetting() {
	itemsUpdateCmd.PersistentFlags().IntVarP(&structArgs.ID, "id", "i", 0, Messages["item"]["id"])
	itemsUpdateCmd.PersistentFlags().StringVarP(&structArgs.Content, "content", "c", "", Messages["item"]["content"])
	itemsUpdateCmd.PersistentFlags().IntVarP(&structArgs.Priority, "priority", "r", 0, Messages["item"]["priority"])
	itemsUpdateCmd.PersistentFlags().IntVar(&structArgs.Collapsed, "collapsed", 0, Messages["item"]["collapsed"])
	itemsUpdateCmd.PersistentFlags().IntVar(&structArgs.AssignedByUID, "assigned-by-uid", 0, Messages["item"]["assigned-by-uid"])
	itemsUpdateCmd.PersistentFlags().IntVar(&structArgs.DayOrder, "day-order", 0, Messages["item"]["day-order"])
	itemsUpdateCmd.MarkPersistentFlagRequired("id")
}

func itemsMoveCmdSetting() {
	itemsMoveCmd.PersistentFlags().IntVarP(&structArgs.ID, "id", "i", 0, Messages["item"]["id"])
	itemsMoveCmd.PersistentFlags().IntVarP(&structArgs.ParentID, "parent-id", "P", 0, Messages["item"]["parent-id"])
	itemsMoveCmd.PersistentFlags().IntVarP(&structArgs.ProjectID, "project-id", "p", 0, Messages["item"]["project-id"])
	itemsMoveCmd.MarkPersistentFlagRequired("id")
}

func itemsDeleteCmdSetting() {
	itemsDeleteCmd.PersistentFlags().IntVarP(&structArgs.ID, "id", "i", 0, Messages["item"]["id"])
	itemsDeleteCmd.MarkPersistentFlagRequired("id")
}

func itemsCompleteCmdSetting() {
	itemsCompleteCmd.PersistentFlags().IntVarP(&structArgs.ID, "id", "i", 0, Messages["item"]["id"])
	itemsCompleteCmd.MarkPersistentFlagRequired("id")
}

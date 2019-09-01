package cmd

// Messages is message of flags help.
var Messages = map[string]map[string]string{
	"item": {
		"id":                "The id of the task.",
		"content":           "The text of the task.",
		"project-id":        "The id of the project to add the task to (a number or a temp id). By default the task is added to the user’s Inbox project.",
		"priority":          "The priority of the task (a number between 1 and 4, 4 for very urgent and 1 for natural).",
		"parent-id":         "The id of the parent task. Set to null for root tasks",
		"child-order":       "The order of task. Defines the position of the task among all the tasks with the same parent_id",
		"day-order":         "The order of the task inside the Today or Next 7 days view (a number, where the smallest value would place the task at the top).",
		"collapsed":         "Whether the task’s sub-tasks are collapsed (where 1 is true and 0 is false).",
		"assigned-by-uid":   "The id of user who assigns the current task. This makes sense for shared projects only.\nAccepts 0 or any user id from the list of project collaborators.\nIf this value is unset or invalid, it will be automatically setup to your uid.",
		"auto-reminder":     "When this option is enabled, the default reminder will be added to the new item if it has a due date with time set.\nSee also the auto_reminder user option for more info about the default reminder.",
		"auto-parse-labels": "When this option is enabled, the labels will be parsed from the task content and added to the task.\nIn case the label doesn’t exist, a new one will be created.",
	},
	"project": {
		"id":            "The id of the project.",
		"name":          "The name of the project.",
		"color":         "The color id of the filter, the value between 30 and 49.",
		"parent-id":     "The id of the parent project. Set to null for root projects",
		"child-order":   "The order of project. Defines the position of the task among all the projects with the same parent_id",
		"collapsed":     "Whether the project’s sub-projects are collapsed (where 1 is true and 0 is false).",
		"shared":        "Whether the project is shared (a true or false value).",
		"is-deleted":    "Whether the project is marked as deleted (where 1 is true and 0 is false).",
		"is-archived":   "Whether the project is marked as archived (where 1 is true and 0 is false).",
		"is-favorite":   "Whether the project is favorite (where 1 is true and 0 is false).",
		"inbox-project": "Whether the project is Inbox (true or otherwise this property is not sent).",
		"team-inbox":    "Whether the project is TeamInbox (true or otherwise this property is not sent).",
	},
}

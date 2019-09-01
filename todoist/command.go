package todoist

// Command is an object included in params of request to Todoist.
type Command struct {
	Type   string                 `json:"type,omitempty"`
	TempID string                 `json:"temp_id,omitempty"`
	UUID   string                 `json:"uuid,omitempty"`
	Args   map[string]interface{} `json:"args,omitempty"`
}

// Commands is a list of Command in params of request to Todoist.
type Commands []Command

// CommandOption is a function to set value to Command.
type CommandOption func(*Command)

// commandArgs returns a function, which work to set a args.
func commandArgs(args map[string]interface{}) CommandOption {
	return func(c *Command) {
		c.Args = args
	}
}

// newCommand returns a command, which received arguments set.
func newCommand(ctype, id string, options ...CommandOption) Command {
	command := Command{
		Type:   ctype,
		TempID: id,
		UUID:   id,
	}
	for _, option := range options {
		option(&command)
	}
	return command
}

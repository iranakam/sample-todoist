package todoist

// Params is
type Params struct {
	Token         string   `json:"token,omitempty"`
	SyncToken     string   `json:"sync_token,omitempty"`
	ResourceTypes []string `json:"resource_types,omitempty"`
	Commands      Commands `json:"commands,omitempty"`
}

// ParamsOption is a function to set params to Params.
type ParamsOption func(*Params)

// paramsCommand returns a function
func paramsCommand(command Command) ParamsOption {
	return func(p *Params) {
		p.Commands = Commands{command}
	}
}

// newParams returns a Params, which received arguments set.
func newParams(token string, options ...ParamsOption) Params {
	params := Params{
		Token:         token,
		SyncToken:     "*",
		ResourceTypes: []string{"project"},
	}
	for _, option := range options {
		option(&params)
	}
	return params
}

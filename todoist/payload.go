package todoist

import (
	"encoding/json"
	"net/url"
)

// PayloadOption is a function to set value to Values.
type PayloadOption func(*url.Values)

// payloadCommands returns a function, which work to set commands.
func payloadCommands(commands Commands) PayloadOption {
	return func(v *url.Values) {
		commandsText, _ := json.Marshal(commands)
		v.Add("commands", string(commandsText))
	}
}

// payloadResourceTypes returns a function, which work to set resource_types.
func payloadResourceTypes(resourceTypes []string) PayloadOption {
	return func(v *url.Values) {
		resourceTypesText, _ := json.Marshal(resourceTypes)
		v.Add("resource_types", string(resourceTypesText))
	}
}

// newPayload returns a string, which received arguments set.
func newPayload(token string, options ...PayloadOption) string {
	values := url.Values{}
	resourceTypesText, _ := json.Marshal([]string{"all"})

	values.Add("token", token)
	values.Add("sync_token", "*")
	values.Add("resource_types", string(resourceTypesText))

	for _, option := range options {
		option(&values)
	}
	return values.Encode()
}

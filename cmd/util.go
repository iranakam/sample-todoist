package cmd

import "encoding/json"

func structToMap(structArgs Args) map[string]interface{} {
	mapArgs := map[string]interface{}{}
	encoded, _ := json.Marshal(structArgs)
	json.Unmarshal(encoded, &mapArgs)
	return mapArgs
}

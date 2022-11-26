package common

import "singer/utils"

func TransferMessage(rules map[string]map[string]string, field []string) map[string]string {
	formatRules := map[string]string{}
	for index, item := range rules {
		for key, value := range item {
			if utils.InSlice(field, index) {
				formatRules[index+"."+key] = value
			}
		}
	}
	return formatRules
}

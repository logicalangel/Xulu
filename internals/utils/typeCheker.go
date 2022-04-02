package utils

import "xulu/internals/consts"

func IsWordVerb(word string) bool {
	for verb, _ := range consts.Verbs {
		if verb == word {
			return true
		}
	}

	return false
}

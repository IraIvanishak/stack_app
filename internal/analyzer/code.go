package analyzer

import (
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func Merge(techMap map[string]int) map[string]int {

	result := make(map[string]int)
	keys := make([]string, 0, len(techMap))
	for k := range techMap {
		keys = append(keys, k)
	}

	// Distance threshold to consider technologies as similar
	const threshold = 2

	// Check each technology against all others
	for _, key := range keys {
		added := false
		for resultKey := range result {
			if levenshtein.DistanceForStrings([]rune(strings.ToLower(key)), []rune(strings.ToLower(resultKey)), levenshtein.DefaultOptions) < threshold {
				result[resultKey] += techMap[key]
				added = true
				break
			}
		}
		if !added {
			result[key] = techMap[key]
		}
	}

	return result
}

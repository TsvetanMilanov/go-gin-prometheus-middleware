package middleware

import "sort"

func getSortedKeys(input map[string]string) []string {
	keys := make([]string, len(input))

	i := 0
	for k := range input {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}

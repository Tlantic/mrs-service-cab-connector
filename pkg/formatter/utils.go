package formatter

import (
	"fmt"
	"strings"
)

func addPrefixToMapKeys(prefix string, map1 map[string]string) map[string]string {
	prefixedMap := make(map[string]string, 0)

	for key, value := range map1 {
		newKey := strings.Join([]string{prefix, key}, ".")
		prefixedMap[newKey] = value
	}

	return prefixedMap
}

func mergeMaps(map1 map[string]string, map2 map[string]string) map[string]string {
	mergedMap := make(map[string]string, 0)

	// Add first map
	for key, value := range map1 {
		mergedMap[key] = value
	}

	// Add second map
	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func addPrefixToMapKeys2(prefix string, map1 map[string]string) map[string]string {
	prefixedMap := make(map[string]string, 0)

	for key, value := range map1 {
		newKey := prefix + "." + key
		prefixedMap[newKey] = value
	}

	return prefixedMap
}

func mergeMaps2(map1 map[string]string, map2 map[string]string) map[string]string {
	for key, value := range map2 {
		map1[key] = value
	}

	return map1
}

func mergeMaps3(map1 map[string]string, map2 map[string]string) {
	for key, value := range map2 {
		map1[key] = value
	}
}

func PadChar(s string, n int, r rune) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("invalid length %d", n)
	}
	if len(s) > n {
		return s, nil
	}
	return strings.Repeat(string(r), n-len(s)) + s, nil
}

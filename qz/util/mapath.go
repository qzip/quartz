package util

import (
	"fmt"
	"strings"
)

//NestedMap return a flattened path provided as an string array
func NestedMap(mi map[string]interface{}, path []string) interface{} {
	var ret interface{}
	ret = nil
	subPath := mi
	for i, tok := range path {
		sp, ok := subPath[tok]
		if !ok {
			break
		}
		if i+1 == len(path) {
			ret = sp
		}
		subPath, ok = sp.(map[string]interface{})
		if !ok {
			break
		}

	}

	return ret
}

//FlattenMap return a flattened path (saperated by .)
func FlattenMap(mi map[string]interface{}, spath string) interface{} {
	return NestedMap(mi, strings.Split(spath, "."))
}

//IntefaceToStringArray interface array to string array
func IntefaceToStringArray(in []interface{}) []string {
	if in == nil {
		return nil
	}
	sa := make([]string, len(in))
	for i, e := range in {
		sa[i] = fmt.Sprint(e)
	}
	return sa
}

// StripMap includes only the paths specified.
func StripMap(in map[string]interface{}, spaths []string) (out map[string]interface{}) {
	if in == nil || len(spaths) == 0 {
		return
	}
	out = make(map[string]interface{}, 0)
	for _, spath := range spaths {
		if len(spath) == 0 {
			continue
		}
		val := FlattenMap(in, spath)
		if val == nil {
			continue
		}
		subPaths := strings.Split(spath, ".")
		cur := out // top
		for i, subPath := range subPaths {
			if i+1 == len(subPaths) {
				cur[subPath] = val
				break // inner
			}
			mi, ok := cur[subPath]
			if !ok {
				cur[subPath] = make(map[string]interface{}, 0)
				cur = cur[subPath].(map[string]interface{})
			} else {
				m, ok := mi.(map[string]interface{})
				if !ok {
					break // inner
				} else {
					cur = m
				}
			}

		} // for sub paths
	}
	return
}

// ReMap re-maps source --> dest the in map to out map
func ReMap(in map[string]interface{}, remap map[string]string) (out map[string]interface{}) {
	out = make(map[string]interface{})
	for src, dest := range remap {
		if len(src) == 0 {
			continue
		}
		val := FlattenMap(in, src)
		if val == nil {
			continue
		}
		cur := out // top
		subPaths := strings.Split(dest, ".")

		for i, subPath := range subPaths {
			if i+1 == len(subPaths) {
				cur[subPath] = val
				break // inner
			}
			mi, ok := cur[subPath]
			if !ok {
				cur[subPath] = make(map[string]interface{}, 0)
				cur = cur[subPath].(map[string]interface{})
			} else {
				m, ok := mi.(map[string]interface{})
				if !ok {
					break // inner
				} else {
					cur = m
				}
			}

		} // for sub paths
	}
	return out
}

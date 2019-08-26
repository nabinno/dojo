package main

import (
	"regexp"
	"sort"
)

// InArray ....
func InArray(array1 []string, array2 []string) (rc []string) {
	for _, a := range array1 {
		if r, err := regexp.Compile(a); err == nil {
			for _, b := range array2 {
				if r.MatchString(b) {
					rc = append(rc, a)
				}
			}
		}

	}
	rc = removeDuplicate(rc)
	sort.Strings(rc)
	if rc == nil {
		rc = []string{}
	}
	return
}

func removeDuplicate(args []string) (rc []string) {
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			rc = append(rc, args[i])
		}
	}
	return
}

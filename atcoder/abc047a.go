package main

// FightingOverCandies ....
func FightingOverCandies(a, b, c int) (rc string) {
	mapInt := map[int]bool{a: true, b: true, c: true}
	keys := getKeys(mapInt)

	if mapInt[maxInt(keys)-minInt(keys)] {
		rc = "Yes"
	} else {
		rc = "No"
	}
	return
}

func maxInt(slc []int) (rc int) {
	for i, e := range slc {
		if i == 0 || e >= rc {
			rc = e
		}
	}
	return
}

func minInt(slc []int) (rc int) {
	for i, e := range slc {
		if i == 0 || e < rc {
			rc = e
		}
	}
	return
}

func getKeys(mapInt map[int]bool) (rc []int) {
	for key := range mapInt {
		rc = append(rc, key)
	}
	return
}

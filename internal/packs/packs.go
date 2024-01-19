package packs

import "sort"
import "errors"

func ItemsToPacks(items int, packs []int) (map[int]int, int, error) {
	// sort the packs in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	// create a slice to hold the number of each pack required
	result := make(map[int]int)
	// iterate over the packs
	for _, pack := range packs {
		if pack <= 0 {
			return nil, 0, errors.New("packs must be greater than 0")
		}
		result[pack] = 0
		if items < pack || items <= 0 {
			continue
		}
		for {
			if items < pack || items <= 0 {
				break
			}
			result[pack] = result[pack] + 1
			items = items - pack
		}
	}

	smallest_pack := packs[len(packs)-1]
	if items > 0 {
		result[smallest_pack] = result[smallest_pack] + 1
		items = items - smallest_pack
	}
	// return the result
	return result, 0-items, nil
}

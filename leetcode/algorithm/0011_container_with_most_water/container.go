package container

import "sort"

type IndexedElement struct {
	Value int
	Index int
}

func maxArea(height []int) int {
	sorted := make([]IndexedElement, len(height))
	for i, val := range height {
		sorted[i].Value = val
		sorted[i].Index = i
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value;
	})

	result := 0
	ma := sorted[0].Index
	mi := sorted[0].Index
	for _, el := range sorted {
		result = max(result, (el.Index - mi) * el.Value)
		result = max(result, (ma - el.Index) * el.Value)
		
		ma = max(ma, el.Index)
		mi = min(mi, el.Index)
	}

	return result
}

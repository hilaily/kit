package array

// GroupIt
// Parameter
//  total: total count of the list
//  count: number of element in a group
// Example:
//  GroupIt(23,10) will return [{Start: 0, End: 10}, {Start: 10, End 23}]
func GroupIt(total, count int) []*Range {
	ret := make([]*Range, 0, total/count+1)
	start := 0
	end := 0
	for {
		end = start + count
		if end > total {
			end = total
		}
		ret = append(ret, &Range{Start: start, End: end})
		start = end
		if end >= total {
			break
		}
	}
	return ret
}

type Range struct {
	Start int
	End   int
}

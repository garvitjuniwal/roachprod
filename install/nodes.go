package install

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func allNodes(total int) []int {
	r := make([]int, total)
	for i := range r {
		r[i] = i + 1
	}
	return r
}

func ListNodes(s string, total int) ([]int, error) {
	if s == "all" {
		return allNodes(total), nil
	}

	m := map[int]bool{}
	for _, p := range strings.Split(s, ",") {
		parts := strings.Split(p, "-")
		switch len(parts) {
		case 1:
			i, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}
			m[i] = true
		case 2:
			from, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
			for i := from; i <= to; i++ {
				m[i] = true
			}
		default:
			return nil, fmt.Errorf("unable to parse nodes specification: %s", p)
		}
	}

	r := make([]int, 0, len(m))
	for i := range m {
		r = append(r, i)
	}
	sort.Ints(r)
	return r, nil
}

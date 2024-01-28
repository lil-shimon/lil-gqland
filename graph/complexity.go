package graph

import "github.com/lil-shimon/lil-gqland/internal"

// ComplexityConfig is a configuration for complexity
func ComplexityConfig() internal.ComplexityRoot {
	var c internal.ComplexityRoot

	c.Repository.Issues = func(childComplexity int, after *string, before *string, first *int, last *int) int {
		var count int
		switch {
		case first != nil && last != nil:
			if *first < *last {
				count = *last
			} else {
				count = *first
			}
		case first != nil && last == nil:
			count = *first
		case first == nil && last != nil:
			count = *last
		default:
			count = 1
		}

		return count * childComplexity
	}
	return c
}

package review

import (
)

// query defines the query of the review project
type Query struct {
	Prompts []string
	Keys    []string
}

func NewQuery(prompts []string, keys []string) (*Query, error) {
	// Create and return the object
	return &Query{
		Prompts: prompts,
		Keys:    keys,
	}, nil
}

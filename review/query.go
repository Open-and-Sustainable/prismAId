package review

import (
)

// query defines the query of the review project
type Query struct {
	Prompts []string
	Keys    []string
}

// NewQuery creates and returns a Query instance based on the provided prompts and keys.
// This function initializes a Query with the given inputs, associating the keys with their corresponding prompts.
//
// Arguments:
// - prompts: A slice of strings containing the input prompts for the review.
// - keys: A slice of strings containing the keys associated with the prompts.
//
// Returns:
// - A Query instance with the specified prompts and keys.
// - An error if the creation fails, although the current implementation does not anticipate errors.
func NewQuery(prompts []string, keys []string) (Query, error) {
	// Create and return the object
	return Query{
		Prompts: prompts,
		Keys:    keys,
	}, nil
}

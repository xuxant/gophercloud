package bays

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a bay resource.
func (r commonResult) Extract() (*Bay, error) {
	var s *Bay
	err := r.ExtractInto(&s)
	return s, err
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// Represents a Container Orchestration Engine Bay, i.e. a cluster
type Bay struct {
	// UUID for the bay
	ID string `json:"uuid"`

	// Human-readable name for the bay. Might not be unique.
	Name string `json:"name"`

	// Indicates whether bay is currently operational. Possible values include:
	// CREATE_IN_PROGRESS, CREATE_FAILED, CREATE_COMPLETE, UPDATE_IN_PROGRESS, UPDATE_FAILED, UPDATE_COMPLETE,
	// DELETE_IN_PROGRESS, DELETE_FAILED, DELETE_COMPLETE, RESUME_COMPLETE, RESTORE_COMPLETE, ROLLBACK_COMPLETE,
	// SNAPSHOT_COMPLETE, CHECK_COMPLETE, ADOPT_COMPLETE.
	Status string `json:"status"`

	// The number of nodes in the bay.
	Nodes int `json:"node_count"`

	// The UUID of the baymodel used to generate the bay.
	BayModelID string `json:"baymodel_id"`
}

// BayPage is the page returned by a pager when traversing over a
// collection of bays.
type BayPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of bays has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r BayPage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, nil
}

// IsEmpty checks whether a BayPage struct is empty.
func (r BayPage) IsEmpty() (bool, error) {
	is, err := ExtractBays(r)
	return len(is) == 0, err
}

// ExtractBays accepts a Page struct, specifically a BayPage struct,
// and extracts the elements into a slice of Bay structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractBays(r pagination.Page) ([]Bay, error) {
	var s struct {
		Bays []Bay `json:"bays"`
	}
	err := (r.(BayPage)).ExtractInto(&s)
	return s.Bays, err
}

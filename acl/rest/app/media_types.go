// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "acl-manager": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-security/acl/rest
// --out=$(GOPATH)/src/github.com/JormungandrK/microservice-security/acl/rest
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// ACLPolicy media type (default view)
//
// Identifier: application/jormungandr-acl-policy+json; view=default
type ACLPolicy struct {
	// Actions to match the request against.
	Actions []string `form:"actions,omitempty" json:"actions,omitempty" xml:"actions,omitempty"`
	// Custom conditions
	Conditions []*Condition `form:"conditions,omitempty" json:"conditions,omitempty" xml:"conditions,omitempty"`
	// Policy description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// allow or deny
	Effect *string `form:"effect,omitempty" json:"effect,omitempty" xml:"effect,omitempty"`
	// Policy ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Resources to which this policy applies.
	Resources []string `form:"resources,omitempty" json:"resources,omitempty" xml:"resources,omitempty"`
	// Subjects to match the request against.
	Subjects []string `form:"subjects,omitempty" json:"subjects,omitempty" xml:"subjects,omitempty"`
}

// Validate validates the ACLPolicy media type instance.
func (mt *ACLPolicy) Validate() (err error) {
	for _, e := range mt.Conditions {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

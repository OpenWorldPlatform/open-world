/*
 * OpenWorldPlatform mgmt service
 *
 * OpenWorldPlatform mgmt
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Project struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

// AssertProjectRequired checks if the required fields are not zero-ed
func AssertProjectRequired(obj Project) error {
	return nil
}

// AssertRecurseProjectRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Project (e.g. [][]Project), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseProjectRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aProject, ok := obj.(Project)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertProjectRequired(aProject)
	})
}

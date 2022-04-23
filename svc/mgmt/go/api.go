/*
 * OpenWorldPlatform mgmt service
 *
 * OpenWorldPlatform mgmt
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	AddProject(http.ResponseWriter, *http.Request)
	AddUser(http.ResponseWriter, *http.Request)
	GetProjectById(http.ResponseWriter, *http.Request)
	Index(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	AddProject(context.Context, Project) (ImplResponse, error)
	AddUser(context.Context, User) (ImplResponse, error)
	GetProjectById(context.Context, string) (ImplResponse, error)
	Index(context.Context) (ImplResponse, error)
}

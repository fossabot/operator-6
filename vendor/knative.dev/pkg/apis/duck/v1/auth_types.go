/*
Copyright 2022 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

// AuthStatus is meant to provide the generated service account name
// in the resource status.
type AuthStatus struct {
	// ServiceAccountName is the name of the generated service account
	// used for this components OIDC authentication.
	ServiceAccountName *string `json:"serviceAccountName,omitempty"`

	// ServiceAccountNames is the list of names of the generated service accounts
	// used for this components OIDC authentication. This list can have len() > 1,
	// when the component uses multiple identities (e.g. in case of a Parallel).
	ServiceAccountNames []string `json:"serviceAccountNames,omitempty"`
}

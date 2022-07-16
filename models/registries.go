/*
Copyright 2022 The kubeall.com Authors.

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

package models

type RegistryPing struct {
	AccessKey      string `json:"access_key" yaml:"access_key"`
	CredentialType string `json:"credential_type" yaml:"credential_type"`
	AccessSecret   string `json:"access_secret" yaml:"access_secret"`
	URL            string `json:"url" yaml:"url"`
	Insecure       bool   `json:"insecure" yaml:"insecure"`
	Type           string `json:"type" yaml:"type"`
	ID             int64  `json:"id" yaml:"id"`
}

type RegistryInfo struct {
	Type                     string         `json:"type"`
	Description              string         `json:"description"`
	SupportedResourceTypes   []string       `json:"-"`
	SupportedResourceFilters []*FilterStyle `json:"supported_resource_filters"`
	SupportedTriggers        []string       `json:"supported_triggers"`
}
type FilterStyle struct {
	Type   string   `json:"type"`
	Style  string   `json:"style"`
	Values []string `json:"values,omitempty"`
}

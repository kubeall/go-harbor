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

import "time"

type AuditLog struct {
	ID           int64     `json:"id" yaml:"id"`
	ProjectID    int64     `json:"project_id" yaml:"project_id"`
	Operation    string    `json:"operation"`
	ResourceType string    `json:"resource_type" yaml:"resource_type"`
	Resource     string    `json:"resource"`
	Username     string    `json:"username"`
	OpTime       time.Time `json:"op_time" yaml:"op_time"`
}

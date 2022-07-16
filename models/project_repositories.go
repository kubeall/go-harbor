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

type Repository struct {
	UpdateTime    time.Time `json:"update_time" yaml:"update_time"`
	Description   string    `json:"description" yaml:"description"`
	PullCount     int64     `json:"pull_count" yaml:"pull_count"`
	CreationTime  time.Time `json:"creation_time" yaml:"creation_time"`
	ArtifactCount int64     `json:"artifact_count" yaml:"artifact_count"`
	ProjectId     int64     `json:"project_id" yaml:"project_id"`
	ID            int64     `json:"id" yaml:"id"`
	Name          string    `json:"name" yaml:"name"`
}
type ProjectRepositoryListOptions struct {
	Q        string `json:"q" yaml:"q"`
	Sort     string `json:"sort" yaml:"sort"`
	Page     int64  `json:"page" yaml:"page"`
	PageSize int64  `json:"page_size" yaml:"page_size"`
}

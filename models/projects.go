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

import (
	"time"
)

type ProjectListOptions struct {
	Q          string `json:"q" yaml:"q"`
	Sort       string `json:"sort" yaml:"sort"`
	Owner      string `json:"owner" yaml:"owner"`
	Name       string `json:"name" yaml:"name"`
	Page       int64  `json:"page" yaml:"page"`
	PageSize   int64  `json:"page_size" yaml:"page_size"`
	Public     bool   `json:"public" yaml:"public"`
	WithDetail bool   `json:"with_detail" yaml:"with_detail"`
}

type ProjectScanner struct {
	UUID *string `json:"uuid" yaml:"uuid"`
}
type ProjectReq struct {
	ProjectName  string          `json:"project_name" yaml:"project_name"`
	CVEAllowlist CVEAllowlist    `json:"cve_allowlist" yaml:"cve_allowlist"`
	RegistryID   int64           `json:"registry_id" yaml:"registry_id"`
	Public       bool            `json:"public" yaml:"public"`
	StorageLimit int64           `json:"storage_limit" yaml:"storage_limit"`
	Metadata     ProjectMetadata `json:"metadata" yaml:"metadata"`
}

// ProjectMetadata holds the metadata of a project.
type ProjectMetadata struct {
	EnableContentTrust       string `json:"enable_content_trust" yaml:"enable_content_trust"`
	EnableContentTrustCosign string `json:"enable_content_trust_cosign" yaml:"enable_content_trust_cosign"`
	AutoScan                 string `json:"auto_scan" yaml:"auto_scan"`
	Severity                 string `json:"severity" yaml:"severity"`
	RetentionId              string `json:"retention_id" yaml:"retention_id"`
	PreventVul               string `json:"prevent_vul" yaml:"prevent_vul"`
	Public                   string `json:"public" yaml:"public"`
	ReuseSysCveAllowlist     string `json:"reuse_sys_cve_allowlist" yaml:"reuse_sys_cve_allowlist"`
}

// Project holds the details of a project.
type Project struct {
	ProjectID    int64             `json:"project_id" yaml:"project_id"`
	OwnerID      int               `json:"owner_id" yaml:"owner_id"`
	Name         string            `json:"name" yaml:"name"`
	CreationTime time.Time         `json:"creation_time" yaml:"creation_time"`
	UpdateTime   time.Time         `json:"update_time" yaml:"update_time"`
	Deleted      bool              `json:"deleted" yaml:"deleted"`
	OwnerName    string            `json:"owner_name" yaml:"owner_name"`
	Role         int               `json:"role" yaml:"role"`
	RoleList     []int             `json:"role_list" yaml:"role_list"`
	RepoCount    int64             `json:"repo_count" yaml:"repo_count"`
	ChartCount   uint64            `json:"chart_count" yaml:"chart_count"`
	Metadata     map[string]string `json:"metadata" yaml:"metadata"`
	CVEAllowlist CVEAllowlist      `json:"cve_allowlist" yaml:"cve_allowlist"`
	RegistryID   int64             `json:"registry_id" yaml:"registry_id"`
}

type CVEAllowlist struct {
	ID           int64              `json:"id" yaml:"id"`
	ProjectID    int64              `json:"project_id" yaml:"project_id"`
	ExpiresAt    *int64             `json:"expires_at" yaml:"expires_at"`
	Items        []CVEAllowlistItem `json:"items" yaml:"items"`
	ItemsText    string             `json:"items_text" yaml:"items_text"`
	CreationTime time.Time          `json:"creation_time" yaml:"creation_time"`
	UpdateTime   time.Time          `json:"update_time" yaml:"update_time"`
}

// CVEAllowlistItem defines one item in the CVE allowlist
type CVEAllowlistItem struct {
	CVEID string `json:"cve_id" yaml:"cve_id"`
}
type Registration struct {
	UpdateTime       time.Time `json:"update_time" yaml:"update_time"`
	Vendor           string    `json:"vendor" yaml:"vendor"`
	Description      string    `json:"description" yaml:"description"`
	Auth             string    `json:"auth" yaml:"auth"`
	IsDefault        bool      `json:"is_default" yaml:"is_default"`
	CreateTime       time.Time `json:"create_time" yaml:"create_time"`
	UUID             string    `json:"uuid" yaml:"uuid"`
	Disabled         bool      `json:"disabled" yaml:"disabled"`
	Name             string    `json:"name" yaml:"name"`
	URL              string    `json:"url" yaml:"url"`
	Adapter          string    `json:"adapter" yaml:"adapter"`
	AccessCredential string    `json:"access_credential" yaml:"access_credential"`
	Version          string    `json:"version" yaml:"version"`
	Health           string    `json:"health" yaml:"health"`
	UseInternalAddr  bool      `json:"use_internal_addr" yaml:"use_internal_addr"`
	SkipCertVerify   bool      `json:"skip_cert_verify" yaml:"skip_cert_verify"`
}

type ProjectSummary struct {
	MaintainerCount   int64               `json:"maintainer_count" yaml:"maintainer_count"`
	ProjectAdminCount int64               `json:"project_admin_count" yaml:"project_admin_count"`
	DeveloperCount    int64               `json:"developer_count" yaml:"developer_count"`
	RepoCount         int64               `json:"repo_count" yaml:"repo_count"`
	Quota             ProjectSummaryQuota `json:"quota" yaml:"quota"`
	GuestCount        int64               `json:"guest_count" yaml:"guest_count"`
	LimitedGuestCount int64               `json:"limited_guest_count" yaml:"limited_guest_count"`
	Registry          Registry            `json:"registry" yaml:"registry"`
	ChartCount        int64               `json:"chart_count" yaml:"chart_count"`
}

//todo
type ProjectSummaryQuota struct {
}

//todo
type Registry struct {
}

type ProjectDeletable struct {
	Message   string `json:"message" yaml:"message"`
	Deletable bool   `json:"deletable" yaml:"deletable"`
}

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

package v1

import (
	"fmt"
	"github.com/kubeall/go-harbor/models"

	"github.com/parnurzeal/gorequest"
	"net/http"
)

type ProjectsService struct {
	client *Client
}

func (s *ProjectsService) UpdateProjectScanner(projectNameOrId string, scanner models.ProjectScanner) (
	success bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.PUT, fmt.Sprintf("projects/%s/scanner", projectNameOrId)).
		Send(scanner).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	return
}

func (s *ProjectsService) GetProjectScanner(projectNameOrId string) (
	scanner models.ProjectScanner, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.PUT, fmt.Sprintf("projects/%s/scanner", projectNameOrId)).
		EndStruct(&scanner)
	return
}

func (s *ProjectsService) GetProjectLogs(projectName string) (
	auditLogs []*models.AuditLog, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, fmt.Sprintf("/projects/%s/logs", projectName)).
		EndStruct(auditLogs)
	return
}
func (s *ProjectsService) CheckProjectExist(name string) (exist bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.HEAD, "projects").
		Query(fmt.Sprintf("project_name=%s", name)).End()
	if resp.StatusCode == http.StatusOK {
		exist = true
	}
	return
}
func (s *ProjectsService) CreateProject(p *models.ProjectReq) (success bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.POST, "projects").
		Send(p).
		End()
	if resp.StatusCode == http.StatusCreated {
		success = true
	}
	return
}
func (s *ProjectsService) ListProjects(option models.ProjectListOptions) (
	projects []*models.Project, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, "projects").
		Query(option).
		EndStruct(projects)
	return
}
func (s *ProjectsService) UpdateProject(p *models.ProjectReq) (success bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.PUT, fmt.Sprintf("projects/%s", p.ProjectName)).
		Send(p).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	return
}
func (s *ProjectsService) GetProjectByNameOrID(projectNameOrId string) (
	project models.Project, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, fmt.Sprintf("projects/%s", projectNameOrId)).
		EndStruct(&project)
	return
}
func (s *ProjectsService) DeleteProject(projectNameOrId string) (success bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.DELETE, fmt.Sprintf("projects/%s", projectNameOrId)).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	return
}
func (s *ProjectsService) GetProjectScannerCandidates(projectNameOrId string) (
	results []*models.Registration, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, fmt.Sprintf("projects/%s/scanner/candidates", projectNameOrId)).
		EndStruct(results)
	return
}
func (s *ProjectsService) GetProjectProjectSummary(projectNameOrId string) (
	result models.ProjectSummary, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, fmt.Sprintf("projects/%s/summary", projectNameOrId)).
		EndStruct(&result)
	return
}
func (s *ProjectsService) GetProjectDeletable(projectNameOrId string) (
	project models.ProjectDeletable, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.newRequest(gorequest.GET, fmt.Sprintf("projects/%s/_deletable", projectNameOrId)).EndStruct(&project)
	return
}

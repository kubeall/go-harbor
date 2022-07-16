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

type ProjectRepositoriesService struct {
	client *Client
}

func (s *ProjectRepositoriesService) UpdateProjectRepository(projectName, repositoryName string,
	repository models.Repository) (success bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.PUT, fmt.Sprintf("projects/%s/repositories/%s", projectName, repositoryName)).
		Send(repository).End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	return
}
func (s *ProjectRepositoriesService) GetProjectRepository(projectName, repositoryName string) (
	repository models.Repository, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, fmt.Sprintf("projects/%s/repositories/%s", projectName, repositoryName)).
		EndStruct(&repository)
	return
}
func (s *ProjectRepositoriesService) DeleteProjectRepository(projectName, repositoryName string) (
	success bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.DELETE, fmt.Sprintf("projects/%s/repositories/%s", projectName, repositoryName)).
		End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	return
}
func (s *ProjectRepositoriesService) ListProjectRepositories(projectName, repositoryName string) (
	repository models.Repository, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.POST, fmt.Sprintf("projects/%s/repositories/%s", projectName, repositoryName)).
		EndStruct(&repository)
	return
}

func (s *ProjectRepositoriesService) ListAllRepositories(option models.ProjectRepositoryListOptions) (
	results []*models.Repository, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, "repositories").
		Query(option).
		EndStruct(results)
	return
}

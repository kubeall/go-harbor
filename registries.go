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

type RegistriesService struct {
	client *Client
}

func (s *RegistriesService) RegistryPing(ping models.RegistryPing) (success bool, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.POST, "registry/ping").
		Send(ping).End()
	if resp.StatusCode == http.StatusOK {
		success = true
	}
	return
}
func (s *RegistriesService) ReplicationAdapters() (adapters []string, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, "registry/ping").
		EndStruct(&adapters)
	return
}
func (s *RegistriesService) RegistryInfo(id int64) (info models.RegistryInfo, resp gorequest.Response, errs []error) {
	resp, _, errs = s.client.
		newRequest(gorequest.GET, fmt.Sprintf("registries/%d/info", id)).
		EndStruct(&info)
	return
}

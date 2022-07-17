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
	"github.com/parnurzeal/gorequest"
	"k8s.io/klog/v2"
	"net/url"
	"strings"
)

const (
	apiVersion = "v2.0"
	userAgent  = "kubeall-go-harbor-client"
)

type Client struct {
	client     *gorequest.SuperAgent
	baseURL    *url.URL
	apiVersion string
	UserAgent  string
	// services
	Projects            *ProjectsService
	Artifacts           *ArtifactsService
	Products            *ProductsService
	ProjectRepositories *ProjectRepositoriesService
	Registries          *RegistriesService
}

func NewClientV2(baseUrl, username, password string, options ...ClientOptionFunc) (client *Client, err error) {
	client, err = newClient(baseUrl, username, password, options...)
	if err != nil {
		klog.Error(err)
		return
	}
	client.apiVersion = apiVersion
	return
}
func NewClient(baseUrl, username, password string, options ...ClientOptionFunc) (client *Client, err error) {
	client, err = newClient(baseUrl, username, password, options...)
	if err != nil {
		klog.Error(err)
		return
	}
	return
}
func newClient(baseUrl, username, password string, options ...ClientOptionFunc) (client *Client, err error) {
	client = &Client{UserAgent: userAgent}
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}
	baseURL, err := url.Parse(baseUrl)
	if err != nil {
		return
	}
	client.baseURL = baseURL
	var harborClient *gorequest.SuperAgent
	if harborClient == nil {
		harborClient = gorequest.New()
	}
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(client); err != nil {
			return nil, err
		}
	}
	client.client = harborClient
	client.client.SetBasicAuth(username, password)
	client.Projects = &ProjectsService{client: client}
	client.Artifacts = &ArtifactsService{client: client}
	client.Products = &ProductsService{client: client}
	client.ProjectRepositories = &ProjectRepositoriesService{client: client}
	client.Registries = &RegistriesService{client: client}
	return
}
func (c *Client) newRequest(method, subPath string) *gorequest.SuperAgent {
	var u string
	if c.apiVersion == apiVersion {
		u = c.baseURL.String() + "api/" + "v2.0/" + subPath
	} else {
		u = c.baseURL.String() + "api/" + subPath
	}
	h := c.client.Set("Accept", "application/json")
	if c.UserAgent != "" {
		h.Set("User-Agent", c.UserAgent)
	}
	switch method {
	case gorequest.PUT:
		return c.client.Put(u).Set("Content-Type", "application/json")
	case gorequest.POST:
		return c.client.Post(u).Set("Content-Type", "application/json")
	case gorequest.GET:
		return c.client.Get(u)
	case gorequest.HEAD:
		return c.client.Head(u)
	case gorequest.DELETE:
		return c.client.Delete(u)
	case gorequest.PATCH:
		return c.client.Patch(u)
	case gorequest.OPTIONS:
		return c.client.Options(u)
	default:
		return c.client.Get(u)
	}
}

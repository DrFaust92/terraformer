// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vultr

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/vultr/govultr"
)

type ReservedIPGenerator struct {
	VultrService
}

func (g ReservedIPGenerator) createResources(ipList []govultr.ReservedIP) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, ip := range ipList {
		resources = append(resources, terraform_utils.NewSimpleResource(
			ip.ReservedIPID,
			ip.ReservedIPID,
			"vultr_reserved_ip",
			"vultr",
			[]string{}))
	}
	return resources
}

func (g *ReservedIPGenerator) InitResources() error {
	client := g.generateClient()
	output, err := client.ReservedIP.List(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}

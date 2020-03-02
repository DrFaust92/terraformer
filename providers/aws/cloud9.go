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

package aws

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
)

var cloud9AllowEmptyValues = []string{"tags."}

type Cloud9Generator struct {
	AWSService
}

func (g *Cloud9Generator) createResources(environmentIds []string) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, resourceName := range environmentIds {
		resources = append(resources, terraform_utils.NewSimpleResource(
			resourceName,
			resourceName,
			"aws_cloud9_environment_ec2",
			"aws",
			cloud9AllowEmptyValues))
	}
	return resources
}

func (g *Cloud9Generator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := cloud9.New(config)
	output, err := svc.ListEnvironmentsRequest(&cloud9.ListEnvironmentsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output.EnvironmentIds)
	return nil
}

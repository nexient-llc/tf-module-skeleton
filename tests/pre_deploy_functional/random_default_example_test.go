// Copyright 2022 Nexient LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"path"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

var ApprovedProviders = []string{
	"registry.terraform.io/hashicorp/random",
}

func TestRandomDefaultExample(t *testing.T) {
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, path.Join("..", ".."), ".")
	_ = files.CopyFile(path.Join("..", "..", ".tool-versions"), path.Join(tempTestFolder, ".tool-versions"))
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tempTestFolder,
		PlanFilePath: "terraform.tfplan",
	})

	tfplan := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

	// For each change in resource
	for _, resourceChange := range tfplan.RawPlan.ResourceChanges {

		// The resource is on the approved list of providers
		assert.Contains(t, ApprovedProviders, resourceChange.ProviderName)
	}
}

package test

import (
	"path"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestRandomWithCakeExample(t *testing.T) {
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "../..", "examples/with_cake")
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

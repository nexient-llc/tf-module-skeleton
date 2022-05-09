package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var ApprovedProviders = []string{
	"registry.terraform.io/hashicorp/random",
}

func TestRandomDefaultExample(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../..",
		PlanFilePath: "terraform.tfplan",
	})

	tfplan := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

	// For each change in resource
	for _, resourceChange := range tfplan.RawPlan.ResourceChanges {

		// The resource is on the approved list of providers
		assert.Contains(t, ApprovedProviders, resourceChange.ProviderName)
	}
}

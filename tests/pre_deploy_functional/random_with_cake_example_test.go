package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRandomWithCakeExample(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/with_cake",
		PlanFilePath: "terraform.tfplan",
	})

	tfplan := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

	// For each change in resource
	for _, resourceChange := range tfplan.RawPlan.ResourceChanges {

		// The resource is on the approved list of providers
		assert.Contains(t, ApprovedProviders, resourceChange.ProviderName)
	}
}

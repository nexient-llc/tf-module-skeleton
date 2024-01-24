package common

import (
	"regexp"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/nexient-llc/lcaf-component-terratest-common/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("TestAlwaysSucceeds", func(t *testing.T) {
		assert.Equal(t, "foo", "foo", "Should always be the same!")
		assert.NotEqual(t, "foo", "bar", "Should never be the same!")
	})

	// When cloning the skeleton to a new module, you will need to change the below test
	// to meet your needs and add any new tests that apply to your situation.
	t.Run("TestSkeletonDeployedIsInvokable", func(t *testing.T) {
		output := terraform.Output(t, ctx.TerratestTerraformOptions, "string")

		// Output contains only alphanumeric characters and 🍰
		assert.Regexp(t, regexp.MustCompile("^[A-Za-z🍰0-9]+$"), output)

		// Other tests would go here and can use functions from lcaf-component-terratest-common.
		// Examples (from lambda):
		// functionName := terraform.Output(t, ctx.TerratestTerraformOptions, "function_name")
		// require.NotEmpty(t, functionName, "name of deployed lambda should be set")
		// awsApiLambdaClient := test_helper_lambda.GetAWSApiLambdaClient(t)
		// test_helper_lambda.WaitForLambdaSpinUp(t, awsApiLambdaClient, functionName)
		// test_helper_lambda.TestIsLambdaInvokable(t, awsApiLambdaClient, functionName)
		// test_helper_lambda.TestLambdaTags(t, awsApiLambdaClient, functionName, ctx.TestConfig.(*ThisTFModuleConfig).Tags)
	})
}

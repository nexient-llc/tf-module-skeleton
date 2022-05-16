package test

// Basic imports
import (
	"path"
	"regexp"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TerraTestSuiteWithCake struct {
	suite.Suite
	TerraformOptions *terraform.Options
}

// setup to do before any test runs
func (suite *TerraTestSuiteWithCake) SetupSuite() {
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(suite.T(), "../..", "examples/with_cake")
	_ = files.CopyFile(path.Join("..", "..", ".tool-versions"), path.Join(tempTestFolder, ".tool-versions"))
	suite.TerraformOptions = terraform.WithDefaultRetryableErrors(suite.T(), &terraform.Options{
		TerraformDir: tempTestFolder,
	})
	terraform.InitAndApplyAndIdempotent(suite.T(), suite.TerraformOptions)
}

// TearDownAllSuite has a TearDownSuite method, which will run after all the tests in the suite have been run.
func (suite *TerraTestSuiteWithCake) TearDownSuit() {
	terraform.Destroy(suite.T(), suite.TerraformOptions)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRunSuiteWithCake(t *testing.T) {
	suite.Run(t, new(TerraTestSuiteWithCake))
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *TerraTestSuiteWithCake) TestOutputWithCake() {
	output := terraform.Output(suite.T(), suite.TerraformOptions, "string")

	// Output contains only alphanumeric characters and cakes
	suite.Regexp(regexp.MustCompile("^[A-Za-z0-9🍰]+$"), output)

	// Output contains at least one cake
	suite.Regexp(regexp.MustCompile("🍰"), output)
}

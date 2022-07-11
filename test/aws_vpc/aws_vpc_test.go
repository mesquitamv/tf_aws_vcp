package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	t.Parallel()

	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../",
	})

	// Runs at end of test. Ensure that infra will be destroyed after tests
	defer terraform.Destroy(t, terraformOptions)

	// Runs init and apply of Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Check VPC creation
	output_vpc := terraform.Output(t, terraformOptions, "vpc_cidr_block")
	assert.Equal(t, output_vpc, "10.0.0.0/16")

}

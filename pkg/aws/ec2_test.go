package aws

import (
	"context"
	"fmt"
	"testing"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	awssdk "github.com/aws/aws-sdk-go/aws"
	test_terraform "github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/gruntwork-io/terratest/modules/random"
)

const AWS_DEFAULT_REGION = "us-east-1"

func TestIfEc2Exists(t *testing.T) {
	expectedName := fmt.Sprintf("examples-aws-%s", random.UniqueId())
	tfWorkingDir := "examples/aws/ec2"

	apply, cleanupTerraform := setupTerraform(t, tfWorkingDir, map[string]interface{}{
		"instance_name": expectedName,
	})
	defer cleanupTerraform()
	apply()

	cfg, err := awsconfig.LoadDefaultConfig(context.Background(), awsconfig.WithRegion("us-west-1"))
	if err != nil {
		t.Fatalf("unable to load SDK config, %v", err)
	}
	svc := ec2.NewFromConfig(cfg)

	filters := []types.Filter{
		{
			Name:   awssdk.String("tag:Name"),
			Values: []string{expectedName},
		},
	}

	inputs := &ec2.DescribeInstancesInput{
		Filters: filters,
	}
	resp, err := svc.DescribeInstances(context.Background(), inputs)
	if err != nil {
		t.Fatal(err)
	}
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			t.Logf("Instance ID: %s, State: %s\n", *instance.InstanceId, string(instance.State.Name))
		}
	}
}

func setupTerraform(t *testing.T, directory string, terraformVars map[string]interface{}) (func(), func()) {
	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../../", directory)

	terraformOptions := test_terraform.WithDefaultRetryableErrors(t, &test_terraform.Options{
		TerraformDir: exampleFolder,
		Vars:         terraformVars,
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": AWS_DEFAULT_REGION,
		},
	})

	return func() {
			test_terraform.InitAndApply(t, terraformOptions)
		}, func() {
			test_terraform.Destroy(t, terraformOptions)
		}
}

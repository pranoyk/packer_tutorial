package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/packer"
	"github.com/stretchr/testify/assert"
)


func TestPackerHelloWorldExample(t *testing.T) {
	packerOptions := &packer.Options{
		Template: "../ami-template.json",
	}

	amiID := packer.BuildArtifact(t, packerOptions)
	
	assert.NotEmpty(t, amiID, "AMI ID should not be an empty string.")
}

/*
func TestPackerAMI(t *testing.T) {
    packerOptions := &packer.Options{
        Template: "path/to/your/packer/template.json",
    }

    // Build the AMI using Packer
    amiID := packer.BuildAmi(t, packerOptions)

    // Validate the AMI using Terratest AWS helpers
    awsRegion := "your-aws-region"
    awsInstanceType := "your-instance-type"

    // Create an AWS client session
    awsRegionOptions := &aws.GetRegionOptions{
        Profile: "your-aws-profile",
    }
    awsClient := aws.NewAuthenticatedClient(t, awsRegion, awsRegionOptions)

    // Validate the AMI properties
    amiTags := map[string]string{
        "your-tag-key": "your-tag-value",
    }
    ami := aws.GetAmiById(t, awsClient, amiID, awsRegion)
    aws.AssertAmiHasTags(t, ami.ID, amiTags)
    aws.AssertAmiExists(t, awsClient, ami.ID, awsRegion)
    aws.AssertInstanceTypeForAmi(t, ami.ID, awsInstanceType, awsRegion)
}
*/
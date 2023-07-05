package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/packer"
	"github.com/stretchr/testify/assert"
)


func TestPackerHelloWorldExample(t *testing.T) {
	packerOptions := &packer.Options{
		Template: "../ami-template.json",
	}
    region := "eu-north-1"
	amiID := packer.BuildArtifact(t, packerOptions)
	defer aws.DeleteAmi(t , region, amiID)
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
    awsRegion := "eu-north-1"
    awsInstanceType := "t3.micro"

    // Create an AWS client session
    
    awsRegionOptions := &aws.GetRegionOptions{
        Profile: "your-aws-profile",
    }
    awsClient := aws.NewAuthenticatedSession(t, awsRegion, awsRegionOptions)

    // Validate the AMI properties
    ami := aws.GetAmiById(t, awsClient, amiID, awsRegion)
    aws.AssertAmiExists(t, awsClient, ami.ID, awsRegion)
    aws.AssertInstanceTypeForAmi(t, ami.ID, awsInstanceType, awsRegion)
}
*/
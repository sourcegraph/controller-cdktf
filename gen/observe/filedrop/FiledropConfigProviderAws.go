package filedrop


type FiledropConfigProviderAws struct {
	// The region where the role ARN exists that you will be dropping files to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/filedrop#region Filedrop#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Your IAM role that Observe allows to drop data into the particular filedrop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/filedrop#role_arn Filedrop#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}


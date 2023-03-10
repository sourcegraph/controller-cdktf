package dmss3endpoint


type DmsS3EndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_s3_endpoint#create DmsS3Endpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_s3_endpoint#delete DmsS3Endpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}


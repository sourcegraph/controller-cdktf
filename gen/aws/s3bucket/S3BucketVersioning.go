package s3bucket

type S3BucketVersioning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3_bucket#enabled S3Bucket#enabled}.
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3_bucket#mfa_delete S3Bucket#mfa_delete}.
	MfaDelete any `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}

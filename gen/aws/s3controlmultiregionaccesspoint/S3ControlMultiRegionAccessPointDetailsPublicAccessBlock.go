package s3controlmultiregionaccesspoint

type S3ControlMultiRegionAccessPointDetailsPublicAccessBlock struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3control_multi_region_access_point#block_public_acls S3ControlMultiRegionAccessPoint#block_public_acls}.
	BlockPublicAcls any `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3control_multi_region_access_point#block_public_policy S3ControlMultiRegionAccessPoint#block_public_policy}.
	BlockPublicPolicy any `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3control_multi_region_access_point#ignore_public_acls S3ControlMultiRegionAccessPoint#ignore_public_acls}.
	IgnorePublicAcls any `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3control_multi_region_access_point#restrict_public_buckets S3ControlMultiRegionAccessPoint#restrict_public_buckets}.
	RestrictPublicBuckets any `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

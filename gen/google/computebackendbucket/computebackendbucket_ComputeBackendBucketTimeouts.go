package computebackendbucket


type ComputeBackendBucketTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_bucket#create ComputeBackendBucket#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_bucket#delete ComputeBackendBucket#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_bucket#update ComputeBackendBucket#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

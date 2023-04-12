package googlestoragebucketobject


type GoogleStorageBucketObjectTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_object#create GoogleStorageBucketObject#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_object#delete GoogleStorageBucketObject#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_object#update GoogleStorageBucketObject#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


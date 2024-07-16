package googlestoragebucketaccesscontrol


type GoogleStorageBucketAccessControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket_access_control#create GoogleStorageBucketAccessControl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket_access_control#delete GoogleStorageBucketAccessControl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket_access_control#update GoogleStorageBucketAccessControl#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


package googlestoragebucketobject


type GoogleStorageBucketObjectContextsCustom struct {
	// An individual object context. Context keys and their corresponding values must start with an alphanumeric character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket_object#key GoogleStorageBucketObject#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value associated with this context. This field holds the primary information for the given context key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket_object#value GoogleStorageBucketObject#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}


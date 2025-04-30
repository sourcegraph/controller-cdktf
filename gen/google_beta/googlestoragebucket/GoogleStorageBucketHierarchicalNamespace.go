package googlestoragebucket


type GoogleStorageBucketHierarchicalNamespace struct {
	// Set this field true to organize bucket with logical file system structure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_bucket#enabled GoogleStorageBucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}


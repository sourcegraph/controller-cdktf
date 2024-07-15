package googlestoragebucket


type GoogleStorageBucketVersioning struct {
	// While set to true, versioning is fully enabled for this bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#enabled GoogleStorageBucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}


package googlestoragebucket


type GoogleStorageBucketVersioning struct {
	// While set to true, versioning is fully enabled for this bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#enabled GoogleStorageBucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}


package storagebucket

type StorageBucketVersioning struct {
	// While set to true, versioning is fully enabled for this bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/storage_bucket#enabled StorageBucket#enabled}
	Enabled any `field:"required" json:"enabled" yaml:"enabled"`
}

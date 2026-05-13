package storagebucket


type StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig struct {
	// Whether GMEK is restricted for new objects within the bucket.
	//
	// If FullyRestricted, new objects can't be created using GMEK encryption. If NotRestricted or unset, creation of new objects with GMEK encryption is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_bucket#restriction_mode StorageBucket#restriction_mode}
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
}


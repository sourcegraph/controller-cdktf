package googlestoragebucketobject


type GoogleStorageBucketObjectRetention struct {
	// The object retention mode. Supported values include: "Unlocked", "Locked".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket_object#mode GoogleStorageBucketObject#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Time in RFC 3339 (e.g. 2030-01-01T02:03:04Z) until which object retention protects this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket_object#retain_until_time GoogleStorageBucketObject#retain_until_time}
	RetainUntilTime *string `field:"required" json:"retainUntilTime" yaml:"retainUntilTime"`
}


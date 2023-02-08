package googlestoragebucket


type GoogleStorageBucketLifecycleRuleCondition struct {
	// Minimum age of an object in days to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#age GoogleStorageBucket#age}
	Age *float64 `field:"optional" json:"age" yaml:"age"`
	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#created_before GoogleStorageBucket#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#custom_time_before GoogleStorageBucket#custom_time_before}
	CustomTimeBefore *string `field:"optional" json:"customTimeBefore" yaml:"customTimeBefore"`
	// Number of days elapsed since the user-specified timestamp set on an object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#days_since_custom_time GoogleStorageBucket#days_since_custom_time}
	DaysSinceCustomTime *float64 `field:"optional" json:"daysSinceCustomTime" yaml:"daysSinceCustomTime"`
	// Number of days elapsed since the noncurrent timestamp of an object. This 							condition is relevant only for versioned objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#days_since_noncurrent_time GoogleStorageBucket#days_since_noncurrent_time}
	DaysSinceNoncurrentTime *float64 `field:"optional" json:"daysSinceNoncurrentTime" yaml:"daysSinceNoncurrentTime"`
	// One or more matching name prefixes to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#matches_prefix GoogleStorageBucket#matches_prefix}
	MatchesPrefix *[]*string `field:"optional" json:"matchesPrefix" yaml:"matchesPrefix"`
	// Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#matches_storage_class GoogleStorageBucket#matches_storage_class}
	MatchesStorageClass *[]*string `field:"optional" json:"matchesStorageClass" yaml:"matchesStorageClass"`
	// One or more matching name suffixes to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#matches_suffix GoogleStorageBucket#matches_suffix}
	MatchesSuffix *[]*string `field:"optional" json:"matchesSuffix" yaml:"matchesSuffix"`
	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#noncurrent_time_before GoogleStorageBucket#noncurrent_time_before}
	NoncurrentTimeBefore *string `field:"optional" json:"noncurrentTimeBefore" yaml:"noncurrentTimeBefore"`
	// Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#num_newer_versions GoogleStorageBucket#num_newer_versions}
	NumNewerVersions *float64 `field:"optional" json:"numNewerVersions" yaml:"numNewerVersions"`
	// Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#with_state GoogleStorageBucket#with_state}
	WithState *string `field:"optional" json:"withState" yaml:"withState"`
}


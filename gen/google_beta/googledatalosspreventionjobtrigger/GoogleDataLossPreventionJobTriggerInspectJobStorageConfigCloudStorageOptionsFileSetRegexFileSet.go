package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet struct {
	// The name of a Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#bucket_name GoogleDataLossPreventionJobTrigger#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// A list of regular expressions matching file paths to exclude.
	//
	// All files in the bucket that match at
	// least one of these regular expressions will be excluded from the scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#exclude_regex GoogleDataLossPreventionJobTrigger#exclude_regex}
	ExcludeRegex *[]*string `field:"optional" json:"excludeRegex" yaml:"excludeRegex"`
	// A list of regular expressions matching file paths to include.
	//
	// All files in the bucket
	// that match at least one of these regular expressions will be included in the set of files,
	// except for those that also match an item in excludeRegex. Leaving this field empty will
	// match all files by default (this is equivalent to including .* in the list)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#include_regex GoogleDataLossPreventionJobTrigger#include_regex}
	IncludeRegex *[]*string `field:"optional" json:"includeRegex" yaml:"includeRegex"`
}


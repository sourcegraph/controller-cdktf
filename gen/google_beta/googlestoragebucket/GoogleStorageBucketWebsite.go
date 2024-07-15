package googlestoragebucket


type GoogleStorageBucketWebsite struct {
	// Behaves as the bucket's directory index where missing objects are treated as potential directories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#main_page_suffix GoogleStorageBucket#main_page_suffix}
	MainPageSuffix *string `field:"optional" json:"mainPageSuffix" yaml:"mainPageSuffix"`
	// The custom object to return when a requested resource is not found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#not_found_page GoogleStorageBucket#not_found_page}
	NotFoundPage *string `field:"optional" json:"notFoundPage" yaml:"notFoundPage"`
}


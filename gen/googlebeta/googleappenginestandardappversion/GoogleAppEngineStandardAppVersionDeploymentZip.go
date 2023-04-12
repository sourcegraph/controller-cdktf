package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionDeploymentZip struct {
	// Source URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_standard_app_version#source_url GoogleAppEngineStandardAppVersion#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// files count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_standard_app_version#files_count GoogleAppEngineStandardAppVersion#files_count}
	FilesCount *float64 `field:"optional" json:"filesCount" yaml:"filesCount"`
}

